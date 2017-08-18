pipeline {
    agent any
    environment {
        GOCONFIG_PATH="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go"
        SONARCONFIG_PATH = "/var/jenkins_home/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqube"
    }

    stages {

        stage('Test') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh "pwd"
                    sh "go test"
                }
            }
        }
        stage('Build') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh 'go build main.go'
                    sh 'echo Build complete'
                }
            }
        }
        stage('Deploy') {
            steps {
                tool(name: 'go', type: 'go')
                withEnv(["GOROOT=$GOCONFIG_PATH", "PATH+GO=$GOCONFIG_PATH/bin"]) {
                    sh "git fetch --tags --progress https://github.com/Luwade/jenkins-docker.git +refs/heads/development:refs/remotes/origin/development"
                    sh "git push https://${env.username}:${env.password}@github.com/Luwade/jenkins-docker.git HEAD:master -f"
                    //sh 'git push origin HEAD:master'
                }
            }
        }
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