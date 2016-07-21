'use strict';

const gulp = require('gulp');
const changed = require('gulp-changed');

gulp.task('install', ['install:webrtc', 'install:bootstrap', 'install:jquery', 'install:simplepeer', 'install:shen']);

gulp.task('install:webrtc', function() {
    return gulp.src('node_modules/webrtc-adapter/out/adapter.js')
        .pipe(changed('dist/js/'))
        .pipe(gulp.dest('dist/js/'));
});

gulp.task('install:jquery', function() {
    return gulp.src('node_modules/jquery/dist/jquery.min.*')
        .pipe(changed('dist/js/'))
        .pipe(gulp.dest('dist/js/'));
});

gulp.task('install:simplepeer', function() {
    return gulp.src('node_modules/simple-peer/simplepeer.min.*')
        .pipe(changed('dist/js/'))
        .pipe(gulp.dest('dist/js/'));
})

gulp.task('install:bootstrap', ['install:bootstrap:js', 'install:bootstrap:css', 'install:bootstrap:fonts']);

gulp.task('install:bootstrap:js', function() {
    return gulp.src('node_modules/bootstrap/dist/js/bootstrap.min.js')
        .pipe(changed('dist/js/'))
        .pipe(gulp.dest('dist/js/'));
});

gulp.task('install:bootstrap:css', function() {
    return gulp.src([
        'node_modules/bootstrap/dist/css/bootstrap.min.css*',
        'node_modules/bootstrap/dist/css/bootstrap-theme.min.css*'
    ])
        .pipe(changed('dist/css/'))
        .pipe(gulp.dest('dist/css/'));
});

gulp.task('install:bootstrap:fonts', function() {
    return gulp.src('node_modules/bootstrap/dist/fonts/*')
        .pipe(changed('dist/fonts/'))
        .pipe(gulp.dest('dist/fonts/'));
});

gulp.task('install:shen', function() {
    return gulp.src('client/**/*')
        .pipe(changed('dist/'))
        .pipe(gulp.dest('dist/'));
});

gulp.task('clean', function() {
    del.sync(deletePaths);
});
