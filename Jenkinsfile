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
  stage("Create binaries") {
    docker.image("golang:1.8.3").inside("-v ${pwd()}:${goPath}") {
      sh "env"
      sh "go get -u github.com/alecthomas/gometalinter"
    }
  }
}