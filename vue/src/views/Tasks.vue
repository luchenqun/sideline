<template>
  <div class="container" style="padding:5px 15px">
    <div class="detail-box" style="padding;:0px">
      <div class="detail3" style="padding:0px;">
        <el-row :gutter="0">
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">Employer(Keplr)</div>
              <div v-if="address" class="value">{{balance}}</div>
              <div v-else class="value">Please connect keplr</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">Sideline Module</div>
              <div class="value">{{sidelineBalance}}</div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>

  <div class="container">
    <div class="gva-btn-list">
      <el-button type="primary" :disabled="!address" @click="dialogVisible = true;">Create Task</el-button>
    </div>
    <!-- <el-divider /> -->
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
        <template #default="scope"><span :style="{color:taskColor(scope.row.status)}">{{ formatTaskStatus(scope.row.status) }}</span></template>
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
import { computed, onBeforeMount, onUnmounted, ref, watch } from 'vue'
import { useStore } from 'vuex'
import { ElMessage, ElLoading } from 'element-plus';
import { useRouter } from 'vue-router';
import { formatTaskStatus, taskColor } from '@/utils/sideline';

export default {
  name: 'Tasks',

  setup() {
    // store
    const $s = useStore()
    const router = useRouter();
    const address = computed(() => $s.getters['common/wallet/address'])
    const balance = ref("")
    const sideline = ref("sl1v7kq6jje3eq43k7eeue7vjjn9lpp0fvrmdw9wm")
    const sidelineBalance = ref("")
    const tasks = ref([])
    const form = ref({ creator: address.value })
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
    const denom = "wrmb";

    onBeforeMount(async () => {
      getData()
      form.value.creator = address.value
      tid.value = setInterval(() => {
        const block = $s?.state?.common?.blocks?.blocks?.at(-1)
        blockHeight.value = block?.height || 0
      }, 1000)
    });

    onUnmounted(async () => {
      clearInterval(tid.value)
    });

    watch(() => address.value, async () => {
      form.value.creator = address.value
      if (address.value) {
        const { balance: bal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: address.value }, query: { denom } })
        balance.value = bal.amount + denom;
      }
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryTaskAll', {});
      console.log("QueryTaskAll", reply)
      tasks.value = reply.Task

      if (address.value) {
        const { balance: bal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: address.value }, query: { denom } })
        balance.value = bal.amount + denom;
      }

      const { balance: sbal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: sideline.value }, query: { denom } })
      sidelineBalance.value = sbal.amount + denom;
    }

    const initForm = () => {
      formRef.value.resetFields();
      form.value = { creator: address.value };
    };

    const closeDialog = () => {
      dialogVisible.value = false;
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
      balance,
      sideline,
      sidelineBalance,
      tasks,
      dialogVisible,
      rules,
      form,
      formRef,
      getData,
      initForm,
      closeDialog,
      submitCreateTask,
      toDetailTask,
      formatTaskStatus,
      taskColor
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/style/detail.scss";
</style>
