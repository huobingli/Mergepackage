<template>
  <el-container class="main-container">
    <el-header class="contianer_header">签名制作</el-header>
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
    
  },
  methods: {
    MergePackage() {
        
    }
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
