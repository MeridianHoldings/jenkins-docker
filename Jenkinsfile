node {
    stage "Prepare environment"
        checkout scm
        def environment  = docker.build 'golang:1.8.3'

        environment.inside {
            stage "Go version"
                sh "go version"
        }
    stage "Cleanup"
        deleteDir()
}