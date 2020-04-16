pipeline  {

    agent any

    stages {
         stage('Clone') {
             steps {
                sh 'echo "1.Clone Stage"'
                git url: "http://root:logointo@192.168.0.120/root/member_es.git"
                script {
                    build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
                }
             }
         }

         stage('Test') {
             steps {
                sh 'echo "2.Test Stage"'
            }
         }
         stage('Build') {
             steps {
                sh 'echo "3.Build Docker Image Stage"'
                sh "docker build -t member-es:${build_tag} ."
            }

         }
         stage('Push') {
             steps {
                sh 'echo "4.Push Docker Image Stage"'
                withCredentials([usernamePassword(credentialsId: '88-Harbor-pwd', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
                    sh "docker login http://103.37.141.216:8083 -u ${dockerHubUser} -p ${dockerHubPassword}"
                    sh "docker tag member-es:${build_tag} 103.37.141.216:8083/library/member-es:${build_tag}"
                    sh "docker push 103.37.141.216:8083/library/member-es:${build_tag}"
                }
            }
         }
         stage('YAML') {
             steps {
                sh "sed -i 's/#APP_BUILD_VER/${build_tag}/g'  deployment.yaml"
                // 输出新创建的部署 yaml 文件内容
                sh "cat deployment.yaml"

            }

         }
         stage('Deploy') {
             steps {
                sh 'echo "6. Deploy Stage1"'
                sh 'sudo kubectl  --kubeconfig=/root/.kube/config  apply -f deployment.yaml'
                sh 'echo "7. Deploy Stage test over2020-04-16 01:18:4211"'
             }
         }
    }

}