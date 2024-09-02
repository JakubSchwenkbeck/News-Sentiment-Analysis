from flask import Flask, request, jsonify
from nltk.sentiment.vader import SentimentIntensityAnalyzer
from textblob import TextBlob

app = Flask(__name__)

@app.route('/vader', methods=['POST'])
def vader_sentiment():
    """
    Endpoint for performing sentiment analysis using VADER.

    This endpoint accepts a POST request with a JSON payload containing the text to analyze.
    It uses VADER (Valence Aware Dictionary and sEntiment Reasoner) to compute the sentiment scores,
    which include negative, neutral, positive, and compound scores.

    Request:
        - Method: POST
        - Content-Type: application/json
        - JSON Body: { "text": "<your_text_here>" }

    Response:
        - Content-Type: application/json
        - JSON Body: {
            "neg": <negative_score>,
            "neu": <neutral_score>,
            "pos": <positive_score>,
            "compound": <compound_score>
        }

    Returns:
        - JSON: Sentiment scores computed by VADER.
    """
    text = request.json['text']
    sia = SentimentIntensityAnalyzer()
    scores = sia.polarity_scores(text)
    return jsonify(scores)

@app.route('/textblob', methods=['POST'])
def textblob_sentiment():
    """
    Endpoint for performing sentiment analysis using TextBlob.

    This endpoint accepts a POST request with a JSON payload containing the text to analyze.
    It uses TextBlob to compute sentiment attributes: polarity and subjectivity.

    Request:
        - Method: POST
        - Content-Type: application/json
        - JSON Body: { "text": "<your_text_here>" }

    Response:
        - Content-Type: application/json
        - JSON Body: {
            "polarity": <polarity_score>,
            "subjectivity": <subjectivity_score>
        }

    Returns:
        - JSON: Sentiment attributes computed by TextBlob.
    """
    text = request.json['text']
    blob = TextBlob(text)
    return jsonify({
        'polarity': blob.sentiment.polarity,
        'subjectivity': blob.sentiment.subjectivity
    })

if __name__ == '__main__':
    """
    Runs the Flask application with debug mode enabled.

    This is the entry point of the Flask application. It starts the server in debug mode, 
    which provides useful debugging information and auto-reloads the server on code changes.
    """
    app.run(debug=True)
