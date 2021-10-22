import mysql.connector
import csv

db = mysql.connector.connect(host="localhost",    # your host, usually localhost
                     user="root",         # your username
                     passwd="",  # your password
                     db="cfc")        # name of the data base

csvfile = open("MOCK_DATA.csv", 'r', encoding='utf8')
reader = csv.reader(csvfile, delimiter=',')
cur = db.cursor(prepared=True)

# Use all the SQL you like

next(reader)
# print all the first cell of all the rows
count = 1
for row in reader:
    print(row)
    try:
        query= ("INSERT INTO Person (userId, userName, password, firstName, lastName, email, address, phoneNumber, role) VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s);")
        cur.execute(query, (int(row[0]), row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8]))
    except:
        print("Failed to insert" + row)
db.commit()
db.close()