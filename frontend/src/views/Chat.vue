<template>
  <div>
    <div class="container" id="app">
      <div class="row">
        <div class="col-md-12">
          <div class="page-header">
            <h2 class="text-center">歡迎來到《柳笑塵埃》聊天室</h2>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-md-1"></div>
        <div class="col-md-6">
          <div>聊天内容</div>
          <div class="msg-list" id="msg-list">
            <div class="message" v-for="msg in msglist.chat" v-bind:key="msg.content" :class="{ system: msg.type==0|| msg.type==5,other:msg.type==3&&nickname!==msg.nickname,myself:nickname==msg.nickname} ">
              <span class="content" style="white-space: pre-wrap;">{{msg.content}}</span>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div>當前在線用戶數:<font color="red">2</font>
          </div>
          <div class="user-list">
            <div v-for="msg in msglist.invitelist" v-bind:key="msg.content">
              <p class="content" style="white-space: pre-wrap;">{{ msg.content }}</p><br>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-md-1"></div>
      <div class="col-md-10">
        <div class="user-input">
          <div class="usertip text-center">用戶提示</div>
          <div class="form-inline has-success text-center" style="margin-bottom: 10px;">
            <div class="input-group">
              <span class="input-group-addon">您的昵稱</span>
              <input type="text" class="form-control" v-model="nickname" aria-describedby="inputGroupSuccess1Status">
            </div>
            <input type="submit" class="form-control btn-primary text-center" @click="leaveChatRoom" value="離開聊天室">
            <input type="submit" class="form-control btn-primary text-center" @click="enterChatRoom" value="進入聊天室">
          </div>
          <textarea id="chat-content" v-model="chatmessage" rows="3" class="form-control" placeholder="在此收入聊天內容。ctrl/command+enter 換行，enter 發送"></textarea>&nbsp;
          <input @click="sendMessage" type="button" value="發送(Enter)" class="btn-primary form-control">
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      ws: {},
      nickname: "",
      chatmessage: "",
      msglist: {
        chat: [],
        invitelist: [],
      },
    };
  },
  mounted() {},
  methods: {
    sendMessage() {
      let msg = JSON.stringify({ content: this.chatmessage, type: "3" });

      this.ws.send(msg);
    },
    enterChatRoom() {
      let domain = location.hostname;
      this.ws = new WebSocket(
        "ws://" + domain + ":2023/ws?nickname=" + this.nickname
      );
      this.ws.onerror = function (evt) {
        console.log(evt);
      };
      this.ws.onopen = function () {
        console.log("Already connected");
        // WebSocket 已连接上的回调
      };
      this.ws.onmessage = (evt) => {
        let data = JSON.parse(evt.data);
        let type = data.type;
        switch (type) {
          case 0:
            this.msglist.chat.push(data);
            break;
          case 1:
            this.msglist.chat.push(data);
            break;
          case 2:
            this.msglist.invitelist.push(data);
            break;
          case 3:
            this.msglist.chat.push(data);
            break;
          case 5:

            if (this.msglist.chat.nickname !== this.nickname) {
              this.msglist.chat.push(data);
            }else{
              this.msglist = [];
              this.nickname = "";
              this.ws.close();
            }

            // if (this.msglist.chat.nickname == this.nickname) {
            //   this.msglist = [];
            //   this.nickname = "";
            //   this.ws.close();
            // }

            break;
        }
      };
    },
    leaveChatRoom() {
      let msg = JSON.stringify({
        content: this.nickname + ":用戶已離開",
        type: "5",
      });
      this.ws.send(msg);
      // this.ws.close();
    },
  },
};
</script>
 <style>
.msg-list {
  height: 400px;
  overflow: scroll;
  border: 1px solid #ccc;
  background-color: #f3f3f3;
  display: flex;
  flex-direction: column;
}
.message {
  margin: 15px 5px 5px 5px;
  padding: 5px;
  background-color: #fff;
}
.message {
  align-self: flex-start;
}

.message .meta {
  color: #ccc;
  font-size: 12px;
}

.message .author {
  color: #999;
  font-weight: bold;
}

.myself {
  background-color: #b0e46e !important;
  color: #ccc;
  align-self: flex-end;
}

.myself .meta {
  color: #2b2b2b;
}

.system {
  background-color: #f3f3f3;
  color: #ccc;
  align-self: center;
}

.other {
  background-color: #f3f3f3;
  color: #ccc;
  align-self: start;
}

.user-list {
  padding-left: 10px;
  height: 400px;
  overflow: scroll;
  border: 1px solid #ccc;
  background-color: #f3f3f3;
}
.user-list .user {
  background-color: #fff;
  margin: 5px;
}

.user-input {
  margin: 10px;
}
.usertip {
  color: red;
}
</style>