from nltk.sentiment.vader import SentimentIntensityAnalyzer
import sys


# Initialize the VADER sentiment analyzer
sia = SentimentIntensityAnalyzer()

def analyze_sentiment(text):
    return sia.polarity_scores(text)

if __name__ == "__main__":
    text = sys.argv[1]
    scores = analyze_sentiment(text)
    print(scores)
