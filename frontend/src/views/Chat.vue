<template>
  <div>
    <div class="container" id="app">
      <div class="row">
        <div class="col-md-12">
          <div class="page-header">
            <h2 class="text-center">欢迎来到《Go 语言编程之旅：一起用 Go 做项目》聊天室</h2>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-md-1"></div>
        <div class="col-md-6">
          <div>聊天内容</div>
          <div class="msg-list" id="msg-list">
            <div class="message">
              <div class="meta"><span class="author"></span> at </div>
              <div>
                <span class="content" style="white-space: pre-wrap;"></span>
              </div>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div>当前在线用户数：<font color="red">2</font>
          </div>
          <div class="user-list">
            用户：@江中智 加入时间：2021-08-10
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
              <span class="input-group-addon">您的昵称</span>
              <input type="text" class="form-control" aria-describedby="inputGroupSuccess1Status">
            </div>
            <input type="submit" class="form-control btn-primary text-center" value="离开聊天室">
            <input type="submit" class="form-control btn-primary text-center" value="进入聊天室">
          </div>
          <textarea id="chat-content" rows="3" class="form-control" placeholder="在此收入聊天内容。ctrl/command+enter 换行，enter 发送"></textarea>&nbsp;
          <input type="button" value="发送(Enter)" class="btn-primary form-control">
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  mounted() {
    // let ws = new WebSocket("ws://192.168.1.103:8000/ws");
    // ws.onerror = function (evt) {
    //   console.log(evt);
    // };
    // ws.onopen = function () {
    //   console.log("GGGGG");
    //   // WebSocket 已连接上的回调
    // };
  },
};
</script>
 <style>
        .msg-list { height: 400px; overflow: scroll; border: 1px solid #ccc; background-color: #f3f3f3; display: flex;flex-direction: column; }
        .message { margin: 15px 5px 5px 5px; padding: 5px; background-color: #fff; }
        .message { align-self: flex-start; }
        .message .meta { color: #ccc; font-size: 12px; }
        .message .author { color: #999; font-weight: bold; }
        .myself { background-color: #b0e46e !important; align-self: flex-end; }
        .myself .meta { color: #2b2b2b; }

        .system { background-color: #f3f3f3; color: #ccc; align-self: center; }

        .user-list { padding-left: 10px; height: 400px; overflow: scroll; border: 1px solid #ccc; background-color: #f3f3f3; }
        .user-list .user { background-color: #fff; margin: 5px; }

        .user-input { margin: 10px; }
        .usertip { color: red; }
    </style>