#!/usr/bin/env groovy

// this will start an executor on a Jenkins agent with the docker label
node {
  // Setup variables
  // application name will be used in a few places so create a variable and use string interpolation to use it where needed
  String applicationName = "basic-app"
  // a basic build number so that when we build and push to Artifactory we will not overwrite our previous builds
  String buildNumber = "0.1.${env.BUILD_NUMBER}"
  // Path we will mount the project to for the Docker container
  String goPath = "/home/tomcat/go"

  // Checkout the code from Github, stages allow Jenkins to visualize the different sections of your build steps in the UI
  stage('Checkout from GitHub') {
    // No special needs here, if your projects relys on submodules the checkout step would need to be different
    checkout scm
  }

  // Start a docker container using the golang:1.8.0-alpine image, mount the current directory to the goPath we specified earlier
  stage("Check images") {
    sh "docker ps --all"
    sh "pwd && ls"
  }
  stage("Create binaries") {
    docker.image("golang:1.8.3").inside("-v ${pwd()}:${goPath} --user root") {
      sh "env"
      sh "go get -u github.com/alecthomas/gometalinter"
      sh "cd /go/bin && ls"
      sh "gometalinter --install"
      // sh "gometalinter --checkstyle > report.xml"
      sh "go get github.com/axw/gocov/..."
      sh "go get github.com/AlekSi/gocov-xml"
      sh "go test -coverprofile=cover.out"
      sh "gocov convert cover.out | gocov-xml > coverage.xml"
      sh "go get -u github.com/jstemmer/go-junit-report"
      sh "go test -v ./... | go-junit-report > test.xml"
      sh "ls"
    }
  }
}