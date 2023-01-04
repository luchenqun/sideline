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
              <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
                <div class="data">
                  <div class="key">email</div>
                  <div class="value">{{employer.email}}</div>
                </div>
              </el-col>
              <el-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4">
                <div class="data">
                  <div class="key">task</div>
                  <div class="value">{{employer?.taskIds?.length}}</div>
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="container">
    <div class="detail-box">
      <div class="key">employer introduce</div>
      <v-md-preview :text="employer.introduce" ref="preview"></v-md-preview>
    </div>
  </div>

  <div class="container">
    <div class="gva-btn-list">
      <el-button :disabled="!(walletAddress && walletAddress == address)" type="primary" @click="dialogVisible = true;">Add Task</el-button>
    </div>
    <el-table :data="tasks">
      <el-table-column align="left" show-overflow-tooltip label="title" min-width="100" prop="title">
        <template #default="scope">
          <div class="nowrap">
            <router-link show-overflow-tooltip :to="('/task/'+scope.row.id)">{{scope.row.title}}</router-link>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="left" label="description" show-overflow-tooltip min-width="100" prop="description" />
      <el-table-column align="left" label="employer" show-overflow-tooltip min-width="100" prop="employer" />
      <el-table-column align="left" label="remuneration" show-overflow-tooltip min-width="100" prop="remuneration" />
      <el-table-column align="left" label="deposit" show-overflow-tooltip min-width="100" prop="deposit" />
      <el-table-column align="left" label="collateral" show-overflow-tooltip min-width="100" prop="collateral" />
      <el-table-column align="left" label="deadline" show-overflow-tooltip min-width="100" prop="deadline" />
      <el-table-column align="left" min-width="100" label="status">
        <template #default="scope">{{ formatTaskStatus(scope.row.status) }}</template>
      </el-table-column>
      <el-table-column align="right" label="action" width="88">
        <template #default="scope">
          <el-button icon="tickets" size="small" type="primary" link @click="toDetailTask(scope.row)">DETAIL</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :before-close="closeDialog" title="Add Task">
      <el-alert style="margin:-20px 0px 10px 0px;" :title="'A new block is generated in about 1s, The current block height is ' + blockHeight" effect="dark" :closable="false" type="warning" />
      <el-form ref="formRef" :model="form" :rules="rules" label-width="115px">
        <el-form-item label="creator" prop="creator">
          <el-input v-model="form.creator" disabled autocomplete="off" />
        </el-form-item>
        <el-form-item label="title" prop="title">
          <el-input v-model="form.title" autocomplete="off" placeholder="task title" />
        </el-form-item>
        <el-form-item label="remuneration" prop="remuneration">
          <el-input v-model="form.remuneration" @input="form.remuneration = form.remuneration.replace(/[^\d]/g, '')" placeholder="remuneration for developer">
            <template #append>WRMB</template>
          </el-input>
        </el-form-item>
        <el-form-item label="deposit" prop="deposit">
          <el-input v-model="form.deposit" @input="form.deposit = form.deposit.replace(/[^\d]/g, '')" placeholder="deposit for this task">
            <template #append>WRMB</template>
          </el-input>
        </el-form-item>
        <el-form-item label="collateral" prop="collateral">
          <el-input v-model="form.collateral" @input="form.collateral = form.collateral.replace(/[^\d]/g, '')" placeholder="collateral that developer are required to pay for tasks">
            <template #append>WRMB</template>
          </el-input>
        </el-form-item>
        <el-form-item label="deadline" prop="deadline">
          <el-input v-model="form.deadline" placeholder="the latest block height for task submission" autocomplete="off" />
        </el-form-item>
        <el-form-item label="description" prop="description">
          <v-md-editor v-model="form.description" height="400px"></v-md-editor>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button type="primary" @click="submitCreateTask">OK</el-button>
        </div>
      </template>
    </el-dialog>

  </div>

</template>

<script>
import { computed, onBeforeMount, ref, watch, onUnmounted } from 'vue'
import { useStore } from 'vuex'
import { ElMessage, ElLoading } from 'element-plus';
import { useRoute, useRouter } from 'vue-router';
import { formatTaskStatus } from '@/utils/sideline';

export default {
  name: 'Employer',

  setup() {
    // store
    const $s = useStore()
    const route = useRoute();
    const router = useRouter();
    const { address } = route.params
    const walletAddress = computed(() => $s.getters['common/wallet/address'])
    const tasks = ref([])
    const employer = ref({})
    const preview = ref();
    const init = { title: "php task", description: "simple task", remuneration: "100", deposit: "200", collateral: "0", deadline: "88888888" }
    const form = ref({ creator: address })
    const formRef = ref(null);
    const dialogVisible = ref(false)
    const rules = ref({
      creator: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      title: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      description: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      remuneration: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      deposit: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      collateral: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      deadline: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
    });
    const blockHeight = ref(0)
    const tid = ref(undefined)

    onBeforeMount(async () => {
      getData()
      form.value.creator = walletAddress.value
      tid.value = setInterval(() => {
        const block = $s?.state?.common?.blocks?.blocks?.at(-1)
        blockHeight.value = block?.height || 0
      }, 1000)
    });

    onUnmounted(async () => {
      clearInterval(tid.value)
    });

    watch(() => walletAddress.value, async () => {
      form.value.creator = walletAddress.value
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryEmployer', { params: { index: address } });
      employer.value = reply.employer

      let items = []
      for (const id of reply.employer.taskIds) {
        let reply = await $s.dispatch('sideline.sideline/QueryTask', { params: { id } });
        items.push(reply.Task)
      }
      tasks.value = items
    }

    const initForm = () => {
      formRef.value.resetFields();
      form.value.creator = walletAddress.value;
      dialogVisible.value = false;
    };

    const closeDialog = () => {
      initForm()
    }

    const submitCreateTask = () => {
      formRef.value.validate(async (valid) => {
        const denom = "wrmb"
        const { value } = form
        if (valid && value?.introduce != "") {

          console.log("value", value)
          const loading = ElLoading.service({ lock: true, text: 'create task...' });
          try {
            const fee = [{ denom: "wrmb", amount: "0" }]
            const reply = await $s.dispatch("sideline.sideline/sendMsgCreateTask", { value: { ...value, remuneration: value.remuneration + denom, deposit: value.deposit + denom, collateral: value.collateral + denom }, fee });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({
                type: 'success',
                message: 'create task success',
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
            console.log("create task error", error)
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

    const toDetailTask = (row) => {
      router.push({
        name: 'task',
        params: row,
      });
    }

    return {
      blockHeight,
      address,
      walletAddress,
      tasks,
      preview,
      employer,
      dialogVisible,
      rules,
      form,
      formRef,
      getData,
      initForm,
      closeDialog,
      submitCreateTask,
      formatTaskStatus,
      toDetailTask,
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
