# scratchpay-clinic-service

This project is simple API that allows search in multiple clinic providers and display results from all the available clinics by any of the following:
<ul>
    <li>Clinic Name</li>
    <li>State</li>
    <li>Availability [ex: from:09:00, to:20:00]</li>
</ul>
 
It's a Golang project, it was developed using Echo (https://echo.labstack.com), a high performance, extensible, minimalist Go web framework.

<b>Installation</b>
<p><i>Requirements:</i>
Install Docker</p>

<p>Step 1: Clone the project, navigate to the root directory on your terminal and run <b><i> docker build -t scratchpay-clinics-service .
</i></b> to build the docker image.</p>
<p>Step 2: Run <b><i>docker images</i></b> to see the list of all docker images on the machine, the image of this project should have been created.</p>
<p>Step 3. Run <b><i>docker run -d -p 8090:8090 scratchpay-clinics-service</i></b> to run the generated image in a container. The image runs on port 8090 of the host machine. </p>
<p>It exposes a GET endpoint with the details below:</p>
<b>URL</b>: <a href="http://localhost:8090/api/v1/clinic/search">http://localhost:8090/api/v1/clinic/search</a> <br />
<b>Method</b>: GET <br />
<b>Query Parameters</b>
<ul>
    <li><b>name</b>: Name of clinic (ex: City Vet). Optional </li>
    <li><b>state</b>: State code (ex: CA). Optional</li>
    <li><b>from</b>: Opening time (ex: 08:00). Optional</li>
    <li><b>to</b>: Closing time (ex: 17:00). Optional</li>
    <li><b>page</b>: Page number (ex: 1. Default=1). Optional</li>
    <li><b>limit</b>: Number of items per page (ex: 5. Default=10). Optional</li>
</ul>
<p>Sample URL: <b><i>http://localhost:8090/api/v1/clinic/search?state=FL&from=18:00&to=17:00&page=1&limit=5</i></b></p>

<p>Launch your browser or any REST client application (eg Postman) and point to <b><i>http://localhost:8090/api/v1/clinic/search</i></b> to check real-time result. </p>
