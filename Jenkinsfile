pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
        SONARCONFIG_PATH = "/var/jenkins_home/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqube"
    }

    stages {
        stage('Build Image') {
            steps {
                script {
                    docker_home = tool 'docker'
                }
                sh "whoami"
                sh "${docker_home}/bin/docker --help && whoami"
                sh "${docker_home}/bin/docker build -t my-app ."
                sh "${docker_home}/bin/docker run hello-world"
            }
        }
    }
}