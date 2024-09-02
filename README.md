# News Sentiment Analysis

This project is a hybrid application that leverages GoLang and Python to perform sentiment analysis on news articles. The backend, written in GoLang, handles fetching and filtering news articles based on various criteria such as country, time periods, and more. The sentiment analysis is performed using pre-built models available in Python via a Flask API.

## Table of Contents
- [Project Overview](#project-overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

The goal of this project is to analyze the sentiment of news articles from various sources. It allows users to filter news articles by different countries, time periods, and other criteria. The sentiment analysis is performed using Python libraries such as NLTK (VADER) and TextBlob, which are accessible via a Flask API.

## Features

- Fetch and filter news articles based on country, time period, and more.
- Perform sentiment analysis on news articles using VADER (NLTK) and TextBlob.
- Hybrid architecture combining the efficiency of GoLang with the flexibility of Python.
- Simple and extensible codebase suitable for further customization.

## Technologies Used

- **GoLang**: For handling the backend logic, including fetching and filtering news articles.
- **Python**: For sentiment analysis using the VADER and TextBlob models.
- **Flask**: A micro web framework used to expose Python sentiment analysis as an API.
- **NLTK**: Natural Language Toolkit, used for VADER sentiment analysis.
- **TextBlob**: A Python library for processing textual data.

