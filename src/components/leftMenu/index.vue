<template>
    <div id="menu">
        <div class="menu-title">
            <p>AXC系統商</p>
        </div>
        <el-menu default-active="2" class="el-menu-vertical-demo" text-color="#fff" background-color="#304056" router>
            <el-sub-menu>
                <template #title>
                    <el-icon>
                        <location />
                    </el-icon>
                    <span>範例</span>
                </template>
                <el-menu-item index="/example/example1">範例1</el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="product">
                <template #title>
                    <el-icon>
                        <location />
                    </el-icon>
                    <span>產品</span>
                </template>
                <el-menu-item index="/product/order">訂單</el-menu-item>
            </el-sub-menu>
        </el-menu>
    </div>
</template>

<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
} from '@element-plus/icons-vue'

</script>

<style lang="scss" scoped>
#menu {
    .menu-title{
        width:100%;
        height:50px;
        background:#1F2937;
        display:flex;
        align-items:center;
        color:#fff;
        padding-left:30px;
    }
    .el-menu {
      border-right: 0;
    }

}



</style>
