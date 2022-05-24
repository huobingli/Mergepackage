<template>
  <el-container class="main-container">
    <el-header>金融终端&下单合包</el-header>
    <el-main>
      <el-row>
        <el-col :span="6">金融终端包路径</el-col>
        <el-col :span="12">
          <el-input
            placeholder="请输入金融终端包"
            v-model="jrzd_input"
            clearable
          ></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">下单包路径</el-col>
        <el-col :span="12">
          <el-input placeholder="请输入下单包" v-model="xiadan_input" clearable>
          </el-input>
        </el-col>
      </el-row>
    </el-main>

    <el-footer><el-button @click="MergePackage">打包</el-button></el-footer>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
        jrzd_input:'',
        xiadan_input:''
    };
  },
  mounted() {
  },
  methods: {
    MergePackage () {
        let jrzd_package = this.jrzd_input
        let xd_package = this.xiadan_input

        if (jrzd_package == "") {
            // alert("请输入正确的金融终端包")
            this.$message.error('请输入正确的金融终端包');
        }

        if (xd_package == "") {
            // alert("请输入正确的下单包")
            this.$message.error('请输入正确的下单包');
        }

        var extras = "D://GitProject//Mergepackage//web";
        // var cldPro = require("child_process");
        // const { exec } = require('child_process');

        var cmdStr = "chcp 65001 && " + extras + "/test.bat";
        console.log(cmdStr)
        const _this = this;
        child_process.exec(cmdStr, function (error, stdout, stderr) {
            // _this.showCommonAlert("SUC0002");
            console.log(error)
            console.log(stdout)
            console.log(stderr)
        });

        // const { spawn } = require('child_process');
        // const bat = child_process.spawn('cmd.exe', ['/c', 'D://GitProject//Mergepackage//web//test.bat']);

        // bat.stdout.on('data', (data) => {
        // console.log(data.toString());
        // });
    }
  },
};
</script>

<style>
.main-container {
  width: 600px;
  margin: 0px auto;
}
</style>
