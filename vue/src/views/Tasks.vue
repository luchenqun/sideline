<template>
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
        <template #default="scope">{{ formatTaskStatus(scope.row.status) }}</template>
      </el-table-column>
      <el-table-column align="right" label="action" width="88">
        <template #default="scope">
          <el-button icon="tickets" size="small" type="primary" link @click="toDetailTask(scope.row)">DETAIL</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :before-close="closeDialog" title="Add Task">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="90px">
        <el-form-item label="creator" prop="creator">
          <el-input v-model="form.creator" disabled autocomplete="off" />
        </el-form-item>
        <el-form-item label="name" prop="name">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="avatar" prop="avatar">
          <el-input v-model="form.avatar" placeholder="please input you avatar http url" autocomplete="off" />
        </el-form-item>
        <el-form-item label="email" prop="email">
          <el-input v-model="form.email" autocomplete="off" />
        </el-form-item>
        <el-form-item label="introduce" prop="introduce">
          <v-md-editor v-model="form.introduce" height="400px"></v-md-editor>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button type="primary" @click="submitRegistEmployer">OK</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script>
import { computed, onBeforeMount, ref, watch, reactive } from 'vue'
import { useStore } from 'vuex'
import { ElMessage, ElLoading } from 'element-plus';
import { useRouter } from 'vue-router';
import { formatTaskStatus } from '@/utils/sideline';

export default {
  name: 'Tasks',

  setup() {
    // store
    const $s = useStore()
    const router = useRouter();
    const address = computed(() => $s.getters['common/wallet/address'])
    const tasks = ref([])
    const form = ref({ creator: address.value })
    const formRef = ref(null);
    const dialogVisible = ref(false)
    const rules = ref({
      creator: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      name: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      introduce: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      email: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      avatar: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
    });

    onBeforeMount(async () => {
      getData()
    });

    watch(() => address.value, async () => {
      form.value.creator = address.value
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryTaskAll', {});
      console.log("QueryTaskAll", reply)
      tasks.value = reply.Task
    }

    const initForm = () => {
      formRef.value.resetFields();
      form.value = { creator: address.value };
    };

    const closeDialog = () => {
      dialogVisible.value = false;
      initForm()
    }

    const submitRegistEmployer = () => {
      formRef.value.validate(async (valid) => {

        if (valid && form.value.introduce != "") {
          const { value } = form
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

    const toDetailTask = (row) => {
      router.push({
        name: 'task',
        params: row,
      });
    }

    return {
      address,
      tasks,
      dialogVisible,
      rules,
      form,
      formRef,
      getData,
      initForm,
      closeDialog,
      submitRegistEmployer,
      toDetailTask,
      formatTaskStatus
    }
  }
}
</script>

<style scoped>
</style>
