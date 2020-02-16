import hashlib, binascii, os
import csv

def hash_password(password):
    salt = hashlib.sha256(os.urandom(60)).hexdigest().encode("ascii")
    hashedPass = hashlib.pbkdf2_hmac('sha512', password.encode("utf-8"), salt, 1)
    hashedPass = binascii.hexlify(hashedPass)
    return (salt+hashedPass).decode("ascii")

def verify_password(stored_password, provided_password):
    salt = stored_password[:64]
    stored_password = stored_password[64:]
    pwdhash = hashlib.pbkdf2_hmac('sha512', provided_password.encode('utf-8'), salt.encode('ascii'), 1)
    pwdhash = binascii.hexlify(pwdhash).decode('ascii')
    return pwdhash == stored_password

users=[]

f = open("user.txt", "w")
last = 10000000
for i in range(last):
    password = "pass"+str(i)
    hashedPassword = hash_password(password)
    
    content = str(i)+"\tuser"+str(i)+"\t"+hashedPassword+"\tuser"+str(i)+"\tdefault_profile_picture.png\n"
    f.write(content)
    if(i % 100000 == 0):
        print(i)
f.close()