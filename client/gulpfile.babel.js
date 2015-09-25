var gulp = require('gulp')
var uglify = require('gulp-uglify')
var htmlreplace = require('gulp-html-replace');
var source = require('vinyl-source-stream');
var browserify = require('browserify');
var watchify = require('watchify');
var reactify = require('reactify');
var streamify = require('gulp-streamify');


// REF https://gist.github.com/danharper/3ca2273125f500429945
var sourcemaps = require('gulp-sourcemaps')
var buffer = require('vinyl-buffer')

var babel = require('babelify')

var path = {
  HTML: 'src/index.html',
  MINIFIED_OUT: 'build.min.js',
  OUT: 'build.js',
  DEST: 'dist',
  DEST_BUILD: 'dist/build',
  DEST_SRC: 'dist/src',
  ENTRY_POINT: './src/js/app.js'
};

var copy = function(){
  gulp
    .src(path.HTML)
    .pipe(gulp.dest(path.DEST));
}
gulp.task('copy', copy)

gulp.task('watch', function() {
  gulp.watch(path.HTML, ['copy'])

  var bundler = watchify(browserify(path.ENTRY_POINT, { debug: true }).transform(babel))

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
  })

  copy()
  rebundle()
})

gulp.task('default', ['watch']);

gulp.task('build', function(){
  browserify({
    entries: [path.ENTRY_POINT],
    transform: [reactify]
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
      'js': 'build/' + path.MINIFIED_OUT
    }))
    .pipe(gulp.dest(path.DEST));
})

gulp.task('production', ['correctAssetUrl', 'build'])
