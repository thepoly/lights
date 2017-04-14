#!/usr/bin/env python

from http.server import BaseHTTPRequestHandler, HTTPServer
from os import curdir, sep
import io

# HTTPRequestHandler class
class Lights_RequestHandler(BaseHTTPRequestHandler):

  # GET
  def do_GET(self):
        # Send response status code
        self.send_response(200)

        # Send headers
        self.send_header('Content-type','text/html')
        self.end_headers()

        # Send message back to client
        f = open('index.html');
        #print(f.read())
        self.wfile.write(bytes(f.read(),"utf8"));
        # Write content as utf-8 data
        #self.wfile.write(bytes(message, "utf8"))
        return

def run():
  print('starting server...')

  # Server settings
  try:
      server_address = ('127.0.0.1', 80)
      httpd = HTTPServer(server_address, Lights_RequestHandler)
      print('running server...')
      httpd.serve_forever()
  except KeyboardInterrupt:
      print("done")

run()
