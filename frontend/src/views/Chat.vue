<template>
  <div>
    <h1>Chat </h1>
  </div>
</template>
<script>
export default {
  mounted() {
    let ws = new WebSocket("ws://192.168.1.103:8000/ws");
    ws.onerror = function (evt) {
      console.log(evt);
    };
    ws.onopen = function () {
        console.log("GGGGG")
      // WebSocket 已连接上的回调
    };
  },
};
</script>
<style lang="">
</style>