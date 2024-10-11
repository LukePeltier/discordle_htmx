import sqlite3
from collections import Counter

def process_strings_and_store_results(db_path):
    # Connect to the SQLite database
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()

    # Fetch strings from the database
    cursor.execute("SELECT id, text FROM messages")
    rows = cursor.fetchall()

    # Flatten all words into a single list
    all_words = [word for _, message in rows for word in message.split()]

    # Calculate word frequencies
    word_freq = Counter(all_words)

    # Identify less frequent words
    threshold = 10
    rare_words = {word for word, freq in word_freq.items() if freq < threshold}


    # Score each string and insert results into the new table
    for row_id, message in rows:
        score = sum(word in rare_words for word in message.split())
        cursor.execute("INSERT INTO message_scores (message_score_message, score) VALUES (?, ?)", (row_id, score))

    # Commit changes and close the connection
    conn.commit()
    conn.close()

# Example usage
db_path = 'discordle.sqlite'
process_strings_and_store_results(db_path)
