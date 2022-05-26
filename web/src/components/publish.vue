<template>
  <el-container class="main-container">
    <el-header>金融终端发布</el-header>
    <el-main>
      <el-row>
        <el-col :span="6">本地包路径</el-col>
        <el-col :span="12">
          <el-input
            placeholder="请输入本地包路径"
            v-model="zip_input"
            clearable
          ></el-input>
        </el-col>
      </el-row>
    </el-main>

    <el-footer><el-button @click="zip_package">打包</el-button></el-footer>
  </el-container>
</template>

<script>
import Axios from "axios";

export default {
  data() {
    return {
      zip_input:""
    };
  },
  mounted() {},
  methods: {
    zip_package() {

    }
  }
};
</script>

<style>
.main-container {
  width: 600px;
  margin: 0px auto;
}
</style>
