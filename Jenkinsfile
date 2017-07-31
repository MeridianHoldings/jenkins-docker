pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/home/vagrant/go"
    }

    stages {
        stage('Build') {
            steps {
                withEnv(["GOROOT=$GOCONFIG_PATH", "GOPATH=$GOCONFIG_PATH"]) {
                    sh 'go build main.go'
                    sh 'echo Build complete'
                }
            }
        }
        stage('Test') {
            steps {
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'go test'
                }
            }
        }
        stage('Deploy') {
            when {
                expression {
                    currentBuild.result == 'SUCCESS'
                }
            }
            steps {
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'git push origin master'
                }
            }
        }
    }
}