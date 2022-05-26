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
import Axios from "axios";

export default {
  data() {
    return {
      jrzd_input: "",
      xiadan_input: "",
      check_ceshi: "",
      check_publish: "",
      check_zip: "",
    };
  },
  mounted() {},
  methods: {
    MergePackage() {
      let jrzd_package = this.jrzd_input
      let xd_package = this.xiadan_input

      if (jrzd_package == "") {
        this.$message.error("请输入正确的金融终端包")
      }

      if (xd_package == "") {
        this.$message.error("请输入正确的下单包")
      }

      if (jrzd_package != "" && xd_package != "") {
        // 这里localhost发布需要修改成对应机器地址
        let request = "http://localhost:7001/CallFunc?jrzd=" + jrzd_package + "&xiadan=" + xd_package
        console.log(request)
        Axios.get(request, {}).then((res) => {
          console.log(res)
        });
      }
    },
  },
};
</script>

<style>
.main-container {
  width: 600px;
  margin: 0px auto;
}
</style>
