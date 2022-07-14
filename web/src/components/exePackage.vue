<template>
  <el-container class="main-container">
    <el-header class="contianer_header">安装包制作</el-header>
    <el-main>
      <el-row class="package_row">
        <el-col :span="6">行情包</el-col>
        <el-col :span="12">
          <el-select v-model="hq_select_value" placeholder="请选择金融终端版本">
            <el-option
              v-for="item in hq_value"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
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
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row class="package_row">
        <el-col :span="6">选择其他券商</el-col>
        <el-col :span="18">
          <el-checkbox-group v-model="checked_qs_list">
            <el-checkbox
              v-for="item in qs_list"
              :label="item.name"
              :key="item.name"
              >{{ item.title }}</el-checkbox
            >
          </el-checkbox-group>
        </el-col>
      </el-row>
    </el-main>
    <el-footer class="container_footer"
      ><el-button @click="MergePackage">制作</el-button></el-footer
    >
  </el-container>
</template>

<script>
import Axios from "axios";
// import JSON from JSON;
const quanshang_list = [
  { title: "财达", name: "caida" },
  { title: "财富", name: "caifu" },
  { title: "财通", name: "caitong" },
  { title: "长江", name: "changjiang" },
  { title: "大同", name: "datong" },
  { title: "德邦", name: "deibang" },
  { title: "东方", name: "dongfang" },
  { title: "东莞", name: "dongguan" },
  { title: "东海", name: "donghai" },
  { title: "东吴", name: "dongwu" },
  { title: "东亚前海", name: "dongyaqianhai" },
  { title: "方正", name: "fangzheng" },
  { title: "恒泰", name: "hengtai" },
  { title: "华创", name: "huahchuang" },
  { title: "华福", name: "huafu" },
  { title: "华泰", name: "huatai" },
  { title: "华鑫顶点", name: "huaxindingdian" },
  { title: "华鑫奇点", name: "huaxinjidian" },
  { title: "开源", name: "kaiyuan" },
  { title: "万联", name: "wanlian" },
  { title: "兴业", name: "xingye" },
  { title: "银河", name: "yinhe" },
  { title: "中山", name: "zhongshan" },
  { title: "中天国富", name: "zhongtianguofu" },
  { title: "中原", name: "zhongyuan" },
  { title: "中泰", name: "zhongtai" },
  { title: "海通", name: "haitong" },
  { title: "万和", name: "wanhe" },
  { title: "万和极速", name: "wanhejisu" },
  { title: "华龙", name: "hualong" },
  { title: "华融", name: "huarong" },
  { title: "天风", name: "tianfeng" },
  { title: "太平洋", name: "taipingyang" },
  { title: "金元", name: "jinyuan" },
  { title: "中信", name: "zhongxin" },
];
export default {
  data() {
    return {
      hq_select_value: "",
      xd_select_value: "",
      qs_select_value: "",
      hq_value: "",
      xd_value: "",
      qs_value: [
        {
          value: "huatai",
          label: "华泰",
        },
        {
          value: "haitong",
          label: "海通",
        },
        {
          value: "other",
          label: "其他券商",
        },
      ],
      checked_qs_list: [],
      qs_list: quanshang_list,
    };
  },
  mounted() {
    // let request = "http://localhost:7001/CallHqPageParse";
    // console.log(request);
    // Axios.get(request, {}).then((res) => {
    //   console.log(res);
    //   var data = res["data"]["data"];
    //   var begin = data.indexOf("[");
    //   var end = data.indexOf("]");
    //   var a = data.substring(begin, end + 1);
    //   console.log(eval(a));
    //   this.hq_value = eval(a);
    // });
  },
  methods: {
    MergePackage() {
      let hq_select_value = this.hq_select_value
      let qs_select_value = this.qs_select_value

      if (hq_select_value == "") {
        this.$message.error("请输入正确的金融终端包");
      }

      if (qs_select_value == "") {
        this.$message.error("请输入正确的下单包");
      }

      if (qs_select_value != "" && hq_select_value != "") {
        // 这里localhost发布需要修改成对应机器地址
        // let request =
        //   "http://10.10.38.32:7001/CallMergePackage?jrzd=" +
        //   jrzd_package +
        //   "&xiadan=" +
        //   xd_package;
        // console.log(request);
        // Axios.get(request, {}).then((res) => {
        //   console.log(res);
        // });

        console.log(hq_select_value)
        console.log(qs_select_value)

        let url = "http://localhost:7001/CallMergePackage2"
        let param = "hq=" + hq_select_value + "&qs=" + qs_select_value + "&broker=" + this.checked_qs_list

        console.log(this.checked_qs_list)
        Axios.post(url, param).then((res) => {
          console.log(res);
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
  padding: 5px;
}

.container_header {
  margin: 0px auto;
}

.container_footer {
  text-align: center;
}
</style>
