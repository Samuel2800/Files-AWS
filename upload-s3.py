import os
import boto3
import tkinter as tk
from tkinter import filedialog

aws_access_key_id = 'key_id'
aws_secret_access_key = 'secret_key'


path = filedialog.askdirectory(initialdir=os.getcwd(), title="Select a dir", )


#create a resource for s3
s3 = boto3.client('s3', aws_access_key_id=aws_access_key_id, aws_secret_access_key=aws_secret_access_key)

#bucket = s3.Bucket("final-project-test-bucket")

for file in os.listdir(path):
    #upload each file
    #bucket.upload_file(Key=file, Filename=f"files/{file}")
    s3.upload_file(f"files/{file}", "final-project-test-bucket", f"uploaded/{file}")