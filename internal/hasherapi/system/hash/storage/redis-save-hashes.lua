local idGenerator = KEYS[1]
local generatedIds = {}
local id
  
for i, sha3Hash in ipairs(ARGV)
do
    id = redis.call("INCR", idGenerator)
    redis.call("SET", "sha3:"..id, sha3Hash)
    generatedIds[i] = id
end

return generatedIds