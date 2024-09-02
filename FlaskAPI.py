from flask import Flask, request, jsonify
from nltk.sentiment.vader import SentimentIntensityAnalyzer
from textblob import TextBlob

app = Flask(__name__)

@app.route('/vader', methods=['POST'])
def vader_sentiment():
    text = request.json['text']
    sia = SentimentIntensityAnalyzer()
    scores = sia.polarity_scores(text)
    return jsonify(scores)

@app.route('/textblob', methods=['POST'])
def textblob_sentiment():
    text = request.json['text']
    blob = TextBlob(text)
    return jsonify({
        'polarity': blob.sentiment.polarity,
        'subjectivity': blob.sentiment.subjectivity
    })

if __name__ == '__main__':
    app.run(debug=True)
