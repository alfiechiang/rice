<template>
   <div id="order">
      <el-table :data="tableData" style="width: 100%">
         <el-table-column prop="date" label="Date" width="180" />
         <el-table-column prop="name" label="Name" width="180" />
         <el-table-column prop="address" label="Address" />
      </el-table>
      <div class="pagination-block">
         <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[10, 20, 30, 40]"
            :small="small" :disabled="disabled" :background="background" layout="total, sizes, prev, pager, next, jumper"
            :total="400" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
      </div>
   </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const currentPage = ref(1)
const pageSize = ref(10)
const background = ref(false)
const disabled = ref(false)
const small = ref(false)

const tableData = [
   {
      date: '2016-05-03',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
   },
   {
      date: '2016-05-02',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
   },
   {
      date: '2016-05-04',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
   },
   {
      date: '2016-05-01',
      name: 'Tom',
      address: 'No. 189, Grove St, Los Angeles',
   },
]

const handleSizeChange = (val: number) => {
   console.log(`${val} items per page`)
}
const handleCurrentChange = (val: number) => {
   console.log(`current page: ${val}`)
}

</script>
<style scope lang="scss" >
#order{
   background: #fff;
   width: 100%;
   height: 100%;
   padding: 20px;
   .pagination-block{
      margin-top: 15px;
      display: flex;
      justify-content: center;
   }

}

</style>