# Profesional Cloud Developer


## Sources
1. [GCP certification guides for cloud-developer][gcp-cd]  
2. [Exam tips from Josh Laird in medium][medium1]  
3. [Exam tips from Sathish Vj in medium][medium2]
4. [YouTube playlist fom Guy Joe][youtube]


## 简介
首先看youtube视频，总结出有什么样的主题subjects（compute engin, kubernetes等等）。
然后根据各个subjects的关键词，在exam guide里面搜寻相关内容。  
最后我们根据搜索到的相关内容，去阅览相关Google文档，作出概要性的总结。


## 指南简介
考试指南里面有5大块内容（设计高可伸缩，高可用，高可信的应用程序），每一块内容下面

1. Designing highly scalable, available, and reliable cloud-native applications  
   设计高可伸缩，高可用，高可信的应用程序
2. Building and Testing Applications  
   编写和测试程序
3. Deploying applications  
   发布程序
4. Integrating Google Cloud Platform Services  
   结合谷歌云的其他服务共同开发
5. Managing Application Performance Monitoring  
   管理应用程序的性能监控
   
每一大块的内容下面还有1.1，1.2这样的小条目，在1.1，1.2这样的条目下面还有更具体的条目。我们按照主题去搜索我们需要学习的内容。
我们学习内容的最小单元就是这里列举的1.1，1.2这样的条目。
 
 
## App Engine
用`App Engine`这个关键词在GCP指南上搜索，总共有以下三个结果:

1. Implementing appropriate deployment strategies based on the target compute environment (Compute Engine, Google Kubernetes Engine, App Engine). Strategies include:  
   选择合适的部署战略
  - Blue/Green deployments
  - Traffic-splitting deployments
  - Rolling deployments
  - Canary deployments
2. Deploying an application to App Engine. Considerations include:  
  - Scaling configuration
  - Versions
  - Traffic splitting
  - Blue/green deployment
3. Integrating an application with compute services. Tasks include:
  - Implementing service discovery in Google Kubernetes Engine, App Engine, and Compute Engine
  - Writing an application that publishes/consumes from Cloud Pub/Sub
  - Reading instance metadata to obtain application configuration
  - Authenticating users by using Oauth2 Web Flow and Identity Aware Proxy
  - Using the CLI tools
  - Configuring compute services network settings (e.g., subnet, firewall ingress/egress, public/private IPs)

### 选择合适的部署战略
首先了解[App Engine](#App-Engine)中列举的部署战略([参考链接][deploymentStrategies])

- Rolling deployment  
The following diagram shows the deployment pattern: the old version is shown in blue and the new version is shown in green across each server in the cluster.
![diagram_rollDeployment][rollingDeployment]

- Blue/Green, Red/Black or A/B deployment  
  This is another fail-safe process. In this method, two identical production environments work in parallel.  
  One is the currently-running production environment receiving all user traffic (depicted as Blue). The other is a clone of it, but idle (Green). Both use the same database back-end and app configuration.  
  If there is an issue after green becomes live, traffic can be routed back to blue.
  ![diagram_bluegreen1][bluegreen1]
  ![diagram_bluegreen2][bluegreen2]
  Blue-green deployments rely on traffic routing. This can be done by updating DNS CNAMES for hosts. However, long TTL values can delay these changes. Alternatively, you can change the load balancer settings so the changes take effect immediately. Features like connection draining in ELB can be used to serve in-flight connections.
  
- Canary deployment  
Canary deployment is like blue-green, except it’s more risk-averse. Instead of switching from blue to green in one step, you use a phased approach.
![diagram_canary][canary]
The main challenge of canary deployment is to devise a way to route some users to the new application.  
Consider a way to route new users by exploring several techniques:

  - Exposing internal users to the canary deployment before allowing external user access;  
  - Basing routing on the source IP range;  
  - Releasing the application in specific geographic regions;  
  - Using an application logic to unlock new features to specific users and groups. This logic is removed when the application goes live for the rest of the users.

- Traffic-splitting deployment  
When you develop and deploy applications for Google App Engine, like Snapchat or Angry Birds do, you can take advantage of traffic splitting to do your A/B testing or progressive rollout. You can decide to split your user traffic across multipe versions of your application at the same time, via an IP address based splitting logic, via Cookies, or randomly [link](traffic-splitting).

#### 需要掌握的
- 用CLI对App Engine进行流量分割  
	```python
	gcloud app services set-traffic [MY_SERVICE]
	--splits [MY_VERSION1]=[VERSION1_WEIGHT],
	[MY_VERSION2]=[VERSION2_WEIGHT] \
	    --split-by [IP_OR_COOKIE]
	```
	```ruby
require 'redcarpet'
markdown = Redcarpet.new("Hello World!")
puts markdown.to_html
	```


## Other
1. Stackdriver 
1. Endpoints quick start in GCP HP
1. 简单的试题
https://myblockchainexperts.org/gcpfreepracticequestions/
4. 详细的准备笔记Deep drive notes
https://myblockchainexperts.org/2019/06/16/gcp-professional-cloud-developer-deep-dive-notes-preparing-for-the-exam/
https://myblockchainexperts.org/?s=Cloud+developer
5. Best practices for enterprise organizations
https://cloud.google.com/docs/enterprise/best-practices-for-enterprise-organizations
6. Portability and design
Trade off platform specific and portable(Cloud SQL is portable and Cloud Spanner is platform specific)
7. BigQuery  pricing
https://cloud.google.com/bigquery/pricing
8. Choose App engine or Kubernetes engine
Use container or not? can rule out the app engine standard and app engine flexible, k8s engine
Better scale performance may be the point to classify app engine flexible, k8s
Pricing may also one factor(basically app engine is lower except using preemptive vm in gke)

<!-- ## Web links -->
[gcp-cd]: https://cloud.google.com/certification/guides/cloud-developer/
[medium1]: https://medium.com/@joshlaird/google-cloud-certified-professional-cloud-developer-exam-tips-6cbc5ddb9fb8
[medium2]: https://medium.com/@sathishvj/notes-from-my-beta-google-cloud-professional-cloud-developer-exam-e5826f6e5de1
[youtube]: https://www.youtube.com/watch?v=dAlhGSoMQaM&list=PLOYQCApvKhV3vB55UF4pRkq3IUdgTn6IY
[deploymentStrategies]: https://dev.to/mostlyjason/intro-to-deployment-strategies-blue-green-canary-and-more-3a3

<!--## Image links-->
[rollingDeployment]: https://res.cloudinary.com/practicaldev/image/fetch/s--RbA0NHA6--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/divuxihkun2p186c9mye.png "rolling deployment"
[bluegreen1]: https://res.cloudinary.com/practicaldev/image/fetch/s--fJ4tYKdy--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/78dk41w8qmuy9f9pvrf6.png "bluegreen old version"
[bluegreen2]: https://res.cloudinary.com/practicaldev/image/fetch/s--ca9C-wVZ--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/m664yyotixnqncprryf0.png "bulegreen new version"
[canary]: https://res.cloudinary.com/practicaldev/image/fetch/s--7PmOiuG9--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/zvf9rbd1x38umph98zro.png "canary"
[traffic-splitting]: https://googlecloud.tips/tips/008-traffic-splitting-with-app-engine/

<abbr title="Hyper Text Markup Language">HTML</abbr>  

