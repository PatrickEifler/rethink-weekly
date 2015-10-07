import gulp from 'gulp'
import uglify from 'gulp-uglify'
import htmlreplace from 'gulp-html-replace'
import source from 'vinyl-source-stream'
import browserify from 'browserify'
import watchify from 'watchify'
import reactify from 'reactify'
import streamify from 'gulp-streamify'

import sass from 'gulp-sass'

// REF https://gist.github.com/danharper/3ca2273125f500429945
import sourcemaps from 'gulp-sourcemaps'
import buffer from 'vinyl-buffer'
import babel from 'babelify'

import autoprefixer from 'gulp-autoprefixer'

const path = {
  HTML: 'src/index.html',
  MINIFIED_OUT: 'build.min.js',
  OUT: 'build.js',
  DEST: 'dist',
  DEST_BUILD: 'dist/build',
  DEST_SRC: 'dist/src',
  ENTRY_POINT: './src/js/app.js',

  SCSS_ENTRY: 'src/sass/app.scss',
  SCSS_SRC:   'src/sass/**/*.scss',
  CSS_DEST: 'dist/src/css/',

  IMAGE_SRC: 'src/images/**',
  IMAGE_DEST: 'dist/images',
}

const copyImage = function() {
  [path.IMAGE_DEST].map( (imageTo) => {
    gulp.src(path.IMAGE_SRC)
      .pipe(gulp.dest(imageTo))
  })
}
gulp.task('image', copyImage)

const copy = function(){
  gulp
    .src(path.HTML)
    .pipe(gulp.dest(path.DEST))
}
gulp.task('copy', copy)

const style = function() {
  const sassOptions = {
      errLogToConsole: true,
        outputStyle: 'expanded'
  }

  gulp.src(path.SCSS_ENTRY)
    .pipe(sourcemaps.init())
    .pipe(sass(sassOptions).on('error', sass.logError))
    .pipe(sourcemaps.write())
    .pipe(autoprefixer())
    .pipe(gulp.dest(path.CSS_DEST))
}
gulp.task('style', style)

gulp.task('watch', function() {
  gulp.watch(path.HTML, ['copy'])
  gulp.watch(path.SCSS_SRC, ['style'])
  gulp.watch(path.IMAGE_SRC, ['image'])

  const bundler = watchify(browserify(path.ENTRY_POINT, { debug: true, transforms: ["reactify", {"es6": true}] }).transform(babel))

  function rebundle() {
    bundler.bundle()
      .on('error', function(err) { console.error(err); this.emit('end'); })
      .pipe(source(path.OUT))
      .pipe(buffer())
      .pipe(sourcemaps.init({ loadMaps: true }))
      .pipe(sourcemaps.write('./'))
      .pipe(gulp.dest(path.DEST_SRC));
  }

  bundler.on('update', function() {
    console.log('-> bundling...')
    rebundle()
    console.log('[x] Done')
  })

  copy()
  style()
  rebundle()
})

gulp.task('default', ['watch']);

gulp.task('build', function(){
  browserify({
    entries: [path.ENTRY_POINT],
    transform: [babel]
  })
    .bundle()
    .pipe(source(path.MINIFIED_OUT))
    //.pipe(streamify(uglify({file: path.MINIFIED_OUT})))
    .pipe(gulp.dest(path.MINIFIED_OUT))
    .pipe(gulp.dest(path.DEST_BUILD))
})

gulp.task('correctAssetUrl', function(){
  gulp.src(path.HTML)
    .pipe(htmlreplace({
      'js': `build/${path.MINIFIED_OUT}`
    }))
    .pipe(gulp.dest(path.DEST));
})

gulp.task('production', ['correctAssetUrl', 'build'])
