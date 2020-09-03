
---@class ResHeartBeatMessage 
---@field public msgID Protos_Login 
---@field public systime integer 
---@field public offsettime integer 
---@field public timeZone string 
local ResHeartBeatMessage = {}

---@class AttributeInfo 
local AttributeInfo = {}

---@class ResErrorMsgMessage 
---@field public msgID Protos_Login 
---@field public errorId integer 
---@field public errorStr string 
local ResErrorMsgMessage = {}

---@class PlayerBaseInfo 
---@field public playerId integer 
---@field public playerName string 
---@field public sex integer 
---@field public headId integer 
---@field public level integer 
---@field public exp integer 
---@field public vipLevel integer 
local PlayerBaseInfo = {}

---@class ReqHeartBeatMessage 
---@field public msgID Protos_Login 
local ReqHeartBeatMessage = {}

---@class ServerInfo 
---@field public id integer 
---@field public beId integer 
---@field public name string 
local ServerInfo = {}
