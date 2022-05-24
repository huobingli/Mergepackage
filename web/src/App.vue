<template>
  <div id="app">
    <router-view />
  </div>
</template>

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
  font-family: 'Microsoft YaHei', 'Avenir', Helvetica, Arial, sans-serif;
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

/* 两边有箭头的滚动条样式 */
/* ::-webkit-scrollbar {
  width: 12px;
  height: 12px;
}
::-webkit-scrollbar-button,
::-webkit-scrollbar-track-piece {
  background-color: #26282e;
  cursor: pointer;
}
::-webkit-scrollbar-button {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-thumb {
  background: #35373e;
  border-radius: 4px;
  border: 2px solid #26282e;
}
::-webkit-scrollbar-thumb:hover {
  background: #4f4f4f;
}
::-webkit-scrollbar-corner {
  background: #26282e;
}
::-webkit-scrollbar-button:vertical:end,
::-webkit-scrollbar-button:vertical:start,
::-webkit-scrollbar-button:horizontal:start,
::-webkit-scrollbar-button:horizontal:end {
  content: '';
  position: relative;
}
::-webkit-scrollbar-button:vertical:start,
::-webkit-scrollbar-button:vertical:end {
  border-left: 5px solid transparent;
  border-right: 5px solid transparent;
  width: 0;
}
::-webkit-scrollbar-button:horizontal:start,
::-webkit-scrollbar-button:horizontal:end {
  border-top: 5px solid transparent;
  border-bottom: 5px solid transparent;
}
::-webkit-scrollbar-button:vertical:end {
  border-top: 5px solid #35373e;
}
::-webkit-scrollbar-button:vertical:start {
  border-bottom: 5px solid #35373e;
}
::-webkit-scrollbar-button:horizontal:start {
  border-right: 5px solid #35373e;
}
::-webkit-scrollbar-button:horizontal:end {
  border-left: 5px solid #35373e;
} */
</style>
