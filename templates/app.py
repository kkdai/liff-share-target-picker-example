from flask import Flask, request, render_template, abort, redirect, jsonify
from flask_bootstrap import Bootstrap
import uuid
import os
from os.path import join, dirname

app = Flask(__name__, static_folder='static')
bootstrap = Bootstrap(app)

@app.route('/')
def do_get():
    return render_template('index.html')

if __name__ == '__main__':
    app.debug = True
    app.run()
