<template>
  <div class="container">
    <div class="detail-box">
      <div class="detail1">
        <div class="left">
          <el-avatar :src="employer.avatar"> <img src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png" /> </el-avatar>
        </div>
        <div class="right">
          <div class="name">{{employer.name}}</div>
          <div>
            <el-row :gutter="0">
              <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                <div class="data">
                  <div class="key">address</div>
                  <div class="value">{{employer.address}}</div>
                </div>
              </el-col>
              <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                <div class="data">
                  <div class="key">email</div>
                  <div class="value">{{employer.email}}</div>
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
      </div>

      <div class="detail2">
        <el-divider class="detail-divider" />
        <div class="items">
          <div class="item">
            <div class="left">
              <img src="@/assets/project_total.png" alt />
            </div>
            <div class="right">
              <div class="value-special"><span class="c1">{{employer?.taskIds?.length}}</span></div>
              <div class="key">Task Total</div>
            </div>
          </div>
          <el-divider direction="vertical" class="divider" />
          <div class="item">
            <div class="left">
              <img src="@/assets/project_success.png" alt />
            </div>
            <div class="right">
              <div class="value-special"><span class="c2">{{employer?.feedbacks?.length}}</span></div>
              <div class="key">Feedback Total</div>
            </div>
          </div>
        </div>
        <el-divider class="detail-divider" />
      </div>
      <v-md-preview :text="employer.introduce" ref="preview"></v-md-preview>
      <el-divider class="detail-divider" />
    </div>
  </div>

  <div class="container">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button :disabled="employer.status>=3" type="primary" icon="plus" @click="submitAddTask">新建任务</el-button>
      </div>
      <el-table ref="multipleTable" :data="tasks" style="width: 100%" tooltip-effect="dark" row-key="ID">
        <el-table-column align="left" label="任务标识" show-overflow-tooltip prop="taskId">
          <template #default="scope">
            <div class="nowrap">
              <router-link show-overflow-tooltip :to="('/layout/taskDetail/' + route.params.projectId + '/' + scope.row.taskId)">{{scope.row.taskId}}</router-link>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="任务名称" show-overflow-tooltip prop="taskName" />
        <el-table-column align="left" label="创建者" show-overflow-tooltip min-width="100" prop="userName" />
        <el-table-column align="left" label="任务索引" show-overflow-tooltip min-width="100" prop="taskIndex" />
        <el-table-column align="left" label="运行时间" show-overflow-tooltip min-width="100" prop="runDuration">
          <template #default="scope">{{ parseInt(scope.row.runDuration / 1000) + "s" }}</template>
        </el-table-column>
        <el-table-column align="left" label="锁定时间" show-overflow-tooltip width="180">
          <template #default="scope">{{ formatDate(scope.row.lockedTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="发布时间" show-overflow-tooltip width="180">
          <template #default="scope">{{ formatDate(scope.row.publishedTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="上报时间" show-overflow-tooltip width="180">
          <template #default="scope">{{ formatDate(scope.row.resultReportTime) }}</template>
        </el-table-column>
        <el-table-column align="left" show-overflow-tooltip label="结果CID" prop="resultCid">
          <template #default="scope">
            <div class="nowrap">
              <a :href="scope.row.downloadResultCidUrl" target="_blank">{{scope.row.resultCid}}</a>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" fixed="right" label="状态" show-overflow-tooltip prop="status" width="80">
          <template #default="scope"><span :class="[('tstatus' + scope.row.status)]">{{ formatTaskStatus(scope.row.status) }}</span></template>
        </el-table-column>
        <el-table-column align="left" fixed="right" label="操作" width="250">
          <template #default="scope">
            <el-button icon="tickets" size="small" type="primary" link @click="submitDetailTask(scope.row)">详情</el-button>
            <el-button :disabled="scope.row.status != 0" size="small" type="primary" link icon="edit" @click="submitEditTask(scope.row)">编辑</el-button>
            <el-button :disabled="scope.row.status != 0" icon="promotion" size="small" type="primary" link @click="submitPublishTask(scope.row)">发布</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="180">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="small" @click="submitDeleteTask(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button :disabled="scope.row.status != 0" type="primary" link icon="delete" size="small" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { computed, onBeforeMount, ref, watch, reactive } from 'vue'
import { useStore } from 'vuex'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage, ElLoading } from 'element-plus';
import { useRoute } from 'vue-router';

export default {
  name: 'Employer',

  setup() {
    // store
    const $s = useStore()
    const route = useRoute();
    const { address } = route.params
    const walletAddress = computed(() => $s.getters['common/wallet/address'])
    const tasks = ref([])
    const employer = ref({})
    const preview = ref();

    onBeforeMount(async () => {
      getData()
    });

    watch(() => walletAddress.value, async () => {
      console.log('地址改变了，搞一下逻辑呗.........', walletAddress.value)
      employer.value.creator = address.value
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryEmployer', { params: { index: address } });
      employer.value = reply.employer
      // tasks.value = reply.employer
    }

    const initForm = () => {
      formRef.value.resetFields();
      employer.value = { creator: address.value, introduce: '' };
    };

    const closeDialog = () => {
      initForm()
    }

    const submitRegistEmployer = () => {
      formRef.value.validate(async (valid) => {

        if (valid && employer.value.introduce != "") {
          const { value } = employer
          console.log("value", value)
          const loading = ElLoading.service({ lock: true, text: 'registing employer...' });
          try {
            const fee = [{ denom: "wrmb", amount: "20000000" }]
            const reply = await $s.dispatch("sideline.sideline/sendMsgRegistEmployer", { value, fee });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({
                type: 'success',
                message: 'regist employer success',
              })
              await getData()
              closeDialog()
            } else {
              ElMessage({
                type: 'error',
                message: reply.rawLog,
              })
            }
          } catch (error) {
            console.log("regist employer error", error)
            ElMessage({
              type: 'error',
              message: error,
            })
          }
          loading.close()
        } else {
          ElMessage({
            type: 'warning',
            message: 'please check you input',
          })
        }
      });
    }

    return {
      address,
      tasks,
      preview,
      employer,
      Plus,
      getData,
      initForm,
      closeDialog,
      submitRegistEmployer,
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/style/detail.scss";
.gva-search-box {
  .el-form-item {
    margin-bottom: 5px;
  }
}
</style>

<style lang="scss">
.container {
  .github-markdown-body {
    padding: 16px 0px 0px 0px;
    font-size: 16px;
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
      sans-serif, Apple Color Emoji, Segoe UI Emoji;
    line-height: 1.5;
    word-wrap: break-word;
  }
}
</style>
