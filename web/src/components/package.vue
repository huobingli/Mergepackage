<template>
  <el-container class="main-container">
    <el-header>金融终端签名打包</el-header>
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
      <el-row>
        <el-col :span="6">包版本</el-col>
        <el-col :span="12">
          <el-checkbox v-model="check_ceshi">测试包</el-checkbox>
          <el-checkbox v-model="check_publish">发布包</el-checkbox>
          <el-checkbox v-model="check_zip">发布zip包</el-checkbox>
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
