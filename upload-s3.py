import os
import boto3
import tkinter as tk
from tkinter import filedialog


path = filedialog.askdirectory(initialdir=os.getcwd(), title="Select a dir", )


#create a resource for s3
s3 = boto3.resource("s3")

bucket = s3.Bucket("final-project-test-bucket")

for file in os.listdir(path):
    #upload each file
    bucket.upload_file(Key=file, Filename=f"files/{file}")