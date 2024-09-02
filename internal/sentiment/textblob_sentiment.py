from textblob import TextBlob
import sys

def analyze_sentiment(text):
    blob = TextBlob(text)
    return blob.sentiment

if __name__ == "__main__":
    text = sys.argv[1]
    sentiment = analyze_sentiment(text)
    print(sentiment)
