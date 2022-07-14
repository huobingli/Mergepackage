<template>
  <el-container class="main-container">
    <el-header class="contianer_header">测试包制作</el-header>
    <el-main>
      <el-row class="package_row">
        <el-col :span="6">行情包</el-col>
        <el-col :span="12">
          <el-select v-model="hq_select_value" placeholder="请选择金融终端版本" >
            <el-option
                v-for="item in hq_value"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
            </el-select>
        </el-col>
      </el-row>
      <el-row class="package_row">
        <el-col :span="6">下单包</el-col>
        <el-col :span="18">
          <el-select v-model="xd_select_value" placeholder="请选择下单版本">
            <el-option
                v-for="item in xd_value"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
            </el-select>
        </el-col>
      </el-row>
      <el-row class="package_row">
        <el-col :span="6">券商分类</el-col>
        <el-col :span="18">
          <el-select v-model="qs_select_value" placeholder="请选择券商版本">
            <el-option
                v-for="item in qs_value"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
            </el-select>
        </el-col>
      </el-row>
      <el-row class="package_row">
        <el-col :span="6">选择其他券商</el-col>
        <el-col :span="18">
            <el-checkbox-group v-model="qs_list">
                <el-checkbox class="qs_sub_item"  v-for="item in qs_list" :label="item.name" :key="item.name" >{{item.title}}</el-checkbox>
            </el-checkbox-group>
        </el-col>
      </el-row>
    </el-main>
    <el-footer class="container_footer"><el-button @click="MergePackage">制作</el-button></el-footer>
  </el-container>
</template>

<script>
import Axios from "axios";

export default {
  data() {
    return {
        hq_select_value: "",
        xd_select_value: "",
        qs_select_value: "",
        hq_value: "",
        xd_value: "",
        //qs_value: "",
        qs_value: [{
            value: 'huatai',
            label: '华泰'
            }, {
            value: 'haitong',
            label: '海通'
            }, {
            value: 'other',
            label: '其他券商'
            }
        ],
        qs_list:[]
    };
  },
  mounted() {
    let requesthq = "http://localhost:7001/CallHqPageParse";
    console.log(requesthq);
    Axios.get(requesthq, {}).then((res) => {
      console.log(res);
      var data = res["data"]["data"];
      var begin = data.indexOf("[");
      var end = data.indexOf("]");
      var a = data.substring(begin, end + 1);
      console.log(eval(a));
      this.hq_value = eval(a);
    });

    let requestJy = "http://localhost:7001/CallJyPageParse";
    console.log(requestJy);
    Axios.get(requestJy, {}).then((res) => {
      console.log(res);
      var data = res["data"]["data"];
      var begin = data.indexOf("[");
      var end = data.indexOf("]");
      var a = data.substring(begin, end + 1);
      console.log(eval(a));
      this.xd_value = eval(a);
    });
  },
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
        let request = "http://10.10.38.32:7001/CallMergePackage?jrzd=" + jrzd_package + "&xiadan=" + xd_package
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
.package_row {
    /* line-height: 3 */
    padding:5px;
}

.container_header {
  margin: 0px auto;
}

.container_footer {
  text-align: center;
}
</style>
