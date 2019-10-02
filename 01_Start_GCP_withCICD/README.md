#Start GCP with CICD
Record steps and commands for Continuous Integration and Continous Delivery

[Test.md](./test.md)

## Steps to do CICD in GCP(Google Cloud Platform)
1. [Create a repository](#create-a-repository-link) <br>
2. [12](#12)

## Set a project
gcloud config list
gcloud projects list
gcloud config set project myProject

## Create a repository [link](https://cloud.google.com/source-repositories/docs/quickstart)
1. Create a new repository
   `gcloud source repos create hello-world`
2. Clone the repository to local
   `gcloud source repos clone hello-world`
3. Create a Hello World script
   ```Python
   #!/usr/bin/env python
   
   import webapp2
   
   class MainHandler(webapp2.RequestHandler):
       def get(self):
           self.response.write('Hello world!')
   
   app = webapp2.WSGIApplication([
       ('/', MainHandler)
   ], debug=True)
   ```
4. Create a `app.yaml` file
   ```
    runtime: python27
    api_version: 1
    threadsafe: yes

    handlers:
    - url: .*
      script: main.app

    libraries:
    - name: webapp2
      version: "2.5.2"
   ```
5. Push to Cloud Sorce Repository
   ```
   git add .
   git commit -m "Add Hello World app to Cloud Source Repositories"
   git push origin master
   ```

## Enable some APIS
- sourcerepo.googleapis.com | Cloud Source Repositories API
- appengine.googleapis.com | App Engine Admin API
- cloudbilling.googleapis.com | Cloud Billing API
- cloudbuild.googleapis.com | Cloud Build API
```
gcloud projects list
gcloud config set project YOUR_PROJECT_ID
gcloud services list --available
gcloud services enable SERVICE_NAME
```

## Grant App Engine access to the Cloud Build service account
1. In the Google Cloud Platform console, open the Cloud Build Settings page:
2. Set the status of the App Engine Admin role to Enable.

## Deploy your application
1. From a terminal window, navigate to the directory containing the repository.
   `cd hello-world`
2. Deploy the sample application.
   `gcloud app deploy app.yaml`
3. Verify that your application is running.
   `gcloud app browse`
4. The browser window should now read:
   `Hello world!`

## Create a cloudbuild.yaml file
1. From a terminal window, navigate to the directory containing the repository.
   `cd hello-world`
2. Using a text editor, create a file named cloudbuild.yaml and paste the following:
   ```
    steps:
    - name: "gcr.io/cloud-builders/gcloud"
      args: ["app", "deploy"]
    timeout: "1600s"
   ```

## Add the cloudbuild.yaml file to your repository
1. Add the file to the repository.
   `git add .`
2. Commit the file.
   `git commit -m "Add cloudbuild.yaml file"`
3. Add the contents of the local Git repository to Cloud Source Repositories using the git push command:
   `git push origin master`

## Create a build trigger
1. Open the Cloud Build page in the Google Cloud Platform Console.
2. Select your project and click Open.
3. Click Create trigger.
4. Select Cloud Source Repository.
5. From the list of available repositories, select the hello-world repository, then click Continue.
6. In the Name box, type:
   `App Engine Test`
7. Under Trigger type, select Branch.
8. Under Build Configuration, select cloudbuild.yaml.
9. In the cloudbuild.yaml location box, type:
   `/cloudbuild.yaml`
10. Click Create Trigger.

## Push a change to your application
1. From a terminal window, use a text editor to update the main.py file to read as follows:
   ```Python
    #!/usr/bin/env python

    import webapp2

    class MainHandler(webapp2.RequestHandler):
        def get(self):
    self.response.write('I update automatically!')

    app = webapp2.WSGIApplication([
        ('/', MainHandler)
    ], debug=True)
    ```
2. Add the file to Git.
   `git add .`
3. Commit the file.
   `git commit -m "Update app to demonstrate build triggers"`
4. Add the contents of the local Git repository to Cloud Source Repositories using the git push command:
   `git push origin master`

## View your build in progress
1. Open the Build triggers page in the Google Cloud Platform Console.
2. Select your project and click Open.
3. In the left nav, click Build history.

## Re-test your application
From, a terminal window, open your application:
`gcloud app browse`
The browser window now reads:
`I update automatically!`

## Clean up
### Delete the build trigger
1. Open the Build triggers page in the Google Cloud Platform Console.
2. Select your project and click Open.
3. Locate the trigger you created.
4. Click the More button More button next to the trigger you want to delete.
5. Select Delete.
### Delete the repository
1. In the GCP Console, open the All repositories page for Cloud Source Repositories.
2. Hover over the repository you want to delete and click Settings settings.
3. The General settings page opens.
4. Click Delete this repository.
5. The Remove repository dialog opens.
6. Type the name of the repository you want to delete.
7. Click Delete.
## 12

# Links
- https://cloud.google.com/source-repositories/docs/quickstart
- https://cloud.google.com/sdk/gcloud/reference/services/
- https://cloud.google.com/source-repositories/docs/quickstart-triggering-builds-with-source-repositories
- https://cloud.google.com/sdk/gcloud/reference/config/
