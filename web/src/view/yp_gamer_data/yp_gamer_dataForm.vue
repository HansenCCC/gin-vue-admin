<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="玩家名称:" prop="gamer_username">
          <el-input v-model="formData.gamer_username" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="移动次数:" prop="gamer_move_count">
          <el-input v-model.number="formData.gamer_move_count" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总耗时:" prop="gameover_time">
          <el-date-picker v-model="formData.gameover_time" type="date" placeholder="选择日期"
            :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Yp_gamer_data'
}
</script>

<script setup>
import {
  createYp_gamer_data,
  updateYp_gamer_data,
  findYp_gamer_data
} from '@/api/yp_gamer_data'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  gamer_username: '',
  gamer_move_count: 0,
  gameover_time: new Date(),
})
// 验证规则
const rule = reactive({
  gamer_username: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  gamer_move_count: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  gameover_time: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findYp_gamer_data({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reyp_gamer_data
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createYp_gamer_data(formData.value)
        break
      case 'update':
        res = await updateYp_gamer_data(formData.value)
        break
      default:
        res = await createYp_gamer_data(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style></style>
