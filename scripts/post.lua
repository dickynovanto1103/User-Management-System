math.randomseed(os.time())
math.random(); math.random(); math.random(); --uses this to make random number generator in the request function below functioning well

request = function()
    num = math.random(0,9999999)
    num = math.fmod(num, 200)
    wrk.method = "POST"
    wrk.body = "username=user" .. num .. "&password=pass" .. num
    wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
    return wrk.format(wrk.method, nil, wrk.headers, wrk.body)
end       