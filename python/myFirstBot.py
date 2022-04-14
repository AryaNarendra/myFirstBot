import strike #For timebeing, strike is a private library. Has to be downloaded into the local from https://github.com/Strike-official/python-sdk 
import flask
import requests
from flask import jsonify
from flask import request

app = flask.Flask(__name__)
app.config["DEBUG"] = True

# The public API link of the hosted server has to be added here.
# Use ngrok to easily make your api public
baseAPI=""

@app.route('/', methods=['POST'])
def home():
    ## Create a strike object
    strikeObj = strike.Create("getting_started",baseAPI+"/respondBack")


    # First Question: Whats your name?
    quesObj1 = strikeObj.Question("name").\
                QuestionText().\
                SetTextToQuestion("Hi! What is your name?")
    quesObj1.Answer("true").TextInput()

    # Second Question: Whats your age?            
    quesObj2 = strikeObj.Question("dob").\
                QuestionText().\
                SetTextToQuestion("Hi! What is your date of birth?")
    quesObj2.Answer("true").DateInput()

    return jsonify(strikeObj.Data())



@app.route('/respondBack', methods=['POST'])
def respondBack():
    data = request.get_json()
    name=data["user_session_variables"]["name"]
    dob=data["user_session_variables"]["dob"][0]
    
    
    strikeObj = strike.Create("getting_started", baseAPI)

    question_card = strikeObj.Question("").\
                QuestionText().\
                SetTextToQuestion("Hi! "+name+" You are the choosen few. You are lucky to be born on "+dob+".")

    return jsonify(strikeObj.Data())


app.run(host='0.0.0.0', port=5001)

