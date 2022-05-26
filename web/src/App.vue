<template>
  <div id="app">
    <el-tabs class="app_navigate" v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="行情下单合包" name="home"></el-tab-pane>
      <el-tab-pane label="发布包制作" name="package"></el-tab-pane>
      <el-tab-pane label="发布包上传平台" name="publish"></el-tab-pane>
    </el-tabs>
    <router-view />
  </div>
</template>

<script>

export default {
  data() {
    return {
      activeName: "home",
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(this.activeName)
      this.$router.push({
        name: this.activeName
      })
    },
  },
};
</script>

<style>
html,
body,
div,
span,
table,
table th,
table td,
input {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-family: "Microsoft YaHei", "Avenir", Helvetica, Arial, sans-serif;
}
#app {
  text-align: center;
}
/* 红涨绿跌 */
.red {
  color: #ff3232 !important;
}
.green {
  color: #22ac38 !important;
}
/* 两边是圆弧的滚动条样式 */
.scroll-hover::-webkit-scrollbar,
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
.scroll-hover::-webkit-scrollbar-thumb,
::-webkit-scrollbar-thumb {
  border-radius: 5px;
  height: 30px;
  transition: all 0.4s ease-out;
}

html::-webkit-scrollbar-thumb,
body::-webkit-scrollbar-thumb,
.show-scroll::-webkit-scrollbar-thumb {
  background: #404040 !important;
}

.scroll-hover::-webkit-scrollbar-corner,
::-webkit-scrollbar-corner {
  background-color: rgba(0, 0, 0, 0);
  background: #181818;
}
/* 鼠标移入出现滚动条，移出隐藏效果 */
.scroll-hover {
  overflow: scroll;
}
:hover::-webkit-scrollbar-thumb {
  background: #404040;
}

.app_navigate {
  width: 600px;
  margin: 0px auto;
}
</style>
