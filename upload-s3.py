import os
import boto3

#create a resource for s3
s3 = boto3.resource("s3")

bucket = s3.Bucket("final-project-test-bucket")

for file in os.listdir("files"):
    #upload each file
    bucket.upload_file(Key=file, Filename=f"files/{file}")