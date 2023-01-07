<template>
  <div class="container" style="padding:5px 15px">
    <div class="detail-box" style="padding;:0px">
      <div class="detail3" style="padding:0px;">
        <el-row :gutter="0">
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">Employer</div>
              <div class="value">{{employerBalance}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div v-if="task.developer" class="data">
              <div class="key">Developer</div>
              <div class="value">{{developerBalance}}</div>
            </div>
            <div v-else class="data">
              <div class="key">Keplr</div>
              <div class="value">{{balance}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">Validator</div>
              <div class="value">{{validatorBalance}}</div>
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

  <div class="container" v-if="task.status == 0 || task.status == 1 || task.status == 2 || task.status == 3 || task.status == 4 || task.status == 5 || task.status == 6 || task.status == 7 || task.status == 8">
    <div class="gva-btn-list" style="margin-bottom:0px;">
      <el-button type="primary" v-if="task.status == 0" @click="submitCancelTask">Cancel Task</el-button>
      <el-button type="primary" v-if="task.status == 0" @click="submitDoTask">Apply Task</el-button>
      <el-button type="primary" v-if="task.status == 1 || task.status == 3" @click="dialogVisible=true">Submit Task</el-button>
      <el-button type="primary" v-if="task.status == 2" @click="submitUndoneTask">Undone Task</el-button>
      <el-button type="primary" v-if="task.status == 2 || task.status == 3" @click="submitStartJudgeTask">Judge Task</el-button>
      <el-button type="primary" v-if="task.status == 6" @click="submitVoteTask">Vote Task</el-button>
      <el-button type="primary" v-if="task.status == 6 && false" @click="submitJudgeTask">Judge Task</el-button>
      <el-button type="primary" v-if="task.status == 1 || task.status == 2 || task.status == 3 || task.status == 6" @click="submitFailTask">Fail Task</el-button>
      <el-button type="primary" v-if="task.status == 1 || task.status == 2 || task.status == 3 || task.status == 6" @click="submitSuccessTask">Success Task</el-button>
      <el-button type="primary" v-if="task.status == 4 || task.status == 5 || task.status == 7 || task.status == 8" @click="submitFeedbackTask">Feedback Task</el-button>
    </div>
  </div>
  <div class="container">
    <el-alert v-if="msg" style="margin:0px 0px 15px 0px;" :title="msg" effect="dark" :closable="false" type="warning" />
    <div class="detail-box">
      <div class="detail1">
        <div class="left">
          <img src="@/assets/task.png" alt />
        </div>
        <div class="right">
          <div class="name">{{task.title}}</div>
          <div>
            <el-row :gutter="0">
              <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
                <div class="data">
                  <div class="key">remuneration</div>
                  <div class="value">{{task.remuneration}}</div>
                </div>
              </el-col>
              <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
                <div class="data">
                  <div class="key">deposit</div>
                  <div class="value">{{task.deposit}}</div>
                </div>
              </el-col>
              <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
                <div class="data">
                  <div class="key">collateral</div>
                  <div class="value">{{task.collateral}}</div>
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
      </div>
      <el-divider class="detail-divider" />
      <div class="detail3" style="padding-bottom:0px;">
        <el-row :gutter="0">
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">deadline</div>
              <div class="value">{{task.deadline}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">status</div>
              <div class="value">{{formatTaskStatus(task.status)}}</div>
            </div>
          </el-col>

          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">employer</div>
              <div class="value nowrap">
                <router-link show-overflow-tooltip :to="('/employer/'+task.employer)">{{task.employer}}</router-link>
              </div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">developer</div>
              <div class="value nowrap">
                <router-link v-if="task.developer" show-overflow-tooltip :to="('/developer/'+task.developer)">{{task.developer}}</router-link>
                <span v-else>no developer do the task</span>
              </div>
            </div>
          </el-col>

          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">accuser</div>
              <div class="value nowrap">{{task.accuser}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">judgeHeight</div>
              <div class="value nowrap">{{task.judgeHeight}}</div>
            </div>
          </el-col>

          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">voteYes</div>
              <div class="value nowrap">{{task.voteYes}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">voteNo</div>
              <div class="value nowrap">{{task.voteNo}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">votedAccounts</div>
              <div class="value nowrap">{{JSON.stringify(task.votedAccounts) }}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">deliverHeight</div>
              <div class="value nowrap">{{task.deliverHeight}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
            <div class="data">
              <div class="key">feedbackByDeveloper</div>
              <div class="value nowrap">{{task.feedbackByDeveloper}}</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
            <div class="data">
              <div class="key">feedbackByEmployer</div>
              <div class="value nowrap">{{task.feedbackByEmployer}}</div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>

  <div class="container">
    <div class="detail-box">
      <div class="key">task description</div>
      <v-md-preview :text="task.description" ref="preview"></v-md-preview>
    </div>
  </div>

  <div class="container">
    <div class="detail-box">
      <div class="key">task deliver</div>
      <v-md-preview :text="task.deliver" ref="preview"></v-md-preview>
    </div>
  </div>

  <el-dialog v-model="dialogVisible" :before-close="closeDialog" title="Submit Task">
    <el-alert title="Please ensure that you have completed the task before submit deliver. Otherwise, the employer have rights initiate arbitration. If you fail, you will lose your collateral" type="warning" effect="dark" style="margin:-20px 0px 20px 0px" />
    <el-form ref="formRef" :model="form" :rules="rules" label-width="70px">
      <el-form-item label="deliver" prop="deliver">
        <v-md-editor v-model="form.deliver" height="400px"></v-md-editor>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="closeDialog">Cancel</el-button>
        <el-button type="primary" @click="submitSubmitTask">OK</el-button>
      </div>
    </template>
  </el-dialog>

</template>

<script>
import { computed, onBeforeMount, onUnmounted, ref, watch } from 'vue'
import { useStore } from 'vuex'
import { ElMessage, ElMessageBox, ElLoading } from 'element-plus'
import { useRoute } from 'vue-router';
import { formatTaskStatus } from '@/utils/sideline';

export default {
  name: 'Employer',

  setup() {
    const init = { title: "php task", description: "simple task", remuneration: "100", deposit: "200", collateral: "0", deadline: "88888888" }
    const vnode = { confirmButtonText: 'OK', cancelButtonText: 'Cancel', type: 'warning' }

    const $s = useStore()
    const route = useRoute();
    const { id } = route.params
    const validator = ref("sl1vmuufd2aen7wkeqz7ya26f90l8c62ey9083rax")
    const validatorBalance = ref("")
    const sideline = ref("sl1v7kq6jje3eq43k7eeue7vjjn9lpp0fvrmdw9wm")
    const sidelineBalance = ref("")
    const developerBalance = ref("")
    const employerBalance = ref("")
    const address = computed(() => $s.getters['common/wallet/address'])
    const balance = ref("")
    const task = ref({})
    const preview = ref();
    const form = ref({})
    const formRef = ref(null);
    const dialogVisible = ref(false)
    const rules = ref({
      deliver: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
    });
    const blockHeight = ref(0)
    const tid = ref(undefined)
    const msg = ref("")
    const params = ref({})
    const denom = "wrmb";

    onBeforeMount(async () => {
      getData()
      tid.value = setInterval(() => {
        const block = $s?.state?.common?.blocks?.blocks?.at(-1)
        blockHeight.value = parseInt(block?.height || 0)
        msg.value = `A new block is generated in about 1s, The current block height is ${blockHeight.value}.If the task is successful, ${params.value.validatorCommission}% of the remuneration will be given to the validator as a commission`
        if (task.value.status == 1 || task.value.status == 3) {
          msg.value = msg.value + ` The task must be submitted before generate ${task.value.deadline - blockHeight.value} new blocks.`
        }
        if (task.value.status == 2) {
          msg.value = msg.value + ` After the developer submit task, the number of confirm blocks to employer is ${params.value.minConfirmSubmitHeight}。So there are only have ${parseInt(task.value.deliverHeight) + params.value.minConfirmSubmitHeight - blockHeight.value} blocks left for employer to confirm.`
        }
        if (task.value.status == 6) {
          msg.value = msg.value + ` After the developer or employer start judge task, the number of vote blocks to validators is ${params.value.minConfirmJudgeHeight}。So there are only have ${parseInt(task.value.judgeHeight) + params.value.minConfirmJudgeHeight - blockHeight.value} blocks left for validator to vote.`
        }
      }, 1000)
    });

    onUnmounted(async () => {
      clearInterval(tid.value)
    });

    watch(() => address.value, async () => {
      if (address.value) {
        const { balance: bal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: address.value }, query: { denom } })
        balance.value = bal.amount + denom;
      }
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryTask', { params: { id } });
      task.value = reply.Task
      form.value.deliver = task.value.deliver;

      reply = await $s.dispatch('sideline.sideline/QueryParams', {});
      // params.value = reply.params
      params.value.minConfirmJudgeHeight = parseInt(reply?.params?.minConfirmJudgeHeight || 0)
      params.value.minConfirmSubmitHeight = parseInt(reply?.params?.minConfirmSubmitHeight || 0)
      params.value.validatorCommission = parseInt(reply?.params?.validatorCommission || 0)

      const { balance: vbal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: validator.value }, query: { denom } })
      validatorBalance.value = vbal.amount + denom;

      const { balance: sbal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: sideline.value }, query: { denom } })
      sidelineBalance.value = sbal.amount + denom;

      if (task.value.developer) {
        const { balance: dbal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: task.value.developer }, query: { denom } })
        developerBalance.value = dbal.amount + denom;
      } else if (address.value) {
        const { balance: bal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: address.value }, query: { denom } })
        balance.value = bal.amount + denom;
      }

      const { balance: ebal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: task.value.employer }, query: { denom } })
      employerBalance.value = ebal.amount + denom;
    }

    const initForm = () => {
      formRef.value.resetFields();
      form.value.deliver = task.value.deliver;
      dialogVisible.value = false;
    };

    const closeDialog = () => {
      initForm()
    }

    const submitCreateTask = () => {
      formRef.value.validate(async (valid) => {
        const denom = "wrmb"
        const { value } = form
        if (valid && value?.introduce) {
          console.log("value", value)
          const loading = ElLoading.service({ lock: true, text: 'create task...' });
          try {
            const fee = [{ denom: "wrmb", amount: "0" }]
            const reply = await $s.dispatch("sideline.sideline/sendMsgCreateTask", { value: { ...value, remuneration: value.remuneration + denom, deposit: value.deposit + denom, collateral: value.collateral + denom }, fee });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'create task success', })
              await getData()
              closeDialog()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("create task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        } else {
          ElMessage({ type: 'warning', message: 'please check you input', })
        }
      });
    }

    const submitCancelTask = async () => {
      if (task.value.employer != address.value) {
        ElMessage({
          type: 'warning',
          message: 'You have no permission to cancel this task, connect keplr or switch to right employer',
        })
        return;
      }
      ElMessageBox.confirm(`it will cancel this task and send back you deposit(${task.value.deposit}) and remuneration(${task.value.remuneration}), continue?`, 'Cancel Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'canceling task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgCancelTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'cancel task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("cancel task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'cancel task canceled', })
        })
    }

    const submitDoTask = async () => {
      // if (false) {
      //   ElMessage({
      //     type: 'warning',
      //     message: 'You are not a developer, please register first',
      //   })
      //   return;
      // }
      ElMessageBox.confirm(`it will take you ${task.value.collateral} for collateral, continue?`, 'Apply Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'applying task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgDoTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'apply task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("apply task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'apply task canceled', })
        })
    }

    const submitSubmitTask = async () => {
      formRef.value.validate(async (valid) => {
        const value = {
          creator: address.value,
          id: task.value.id,
          deliver: form.value.deliver
        }
        if (valid && value?.deliver) {
          console.log("value", value)
          const loading = ElLoading.service({ lock: true, text: 'submiting task...' });
          try {
            const fee = [{ denom: "wrmb", amount: "0" }]
            const reply = await $s.dispatch("sideline.sideline/sendMsgSubmitTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'submit task success', })
              await getData()
              closeDialog()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("submit task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        } else {
          ElMessage({ type: 'warning', message: 'please check you input', })
        }
      });
    }

    const submitUndoneTask = async () => {
      if (task.value.employer != address.value) {
        ElMessage({
          type: 'warning',
          message: 'You have no permission to undone this task, connect keplr or switch to right employer',
        })
        return;
      }
      ElMessageBox.confirm(`Please ensure that developer not completed the task before submit undone. Otherwise, the developer have rights initiate arbitration. If you fail, you will lose your deposit(${task.value.deposit}) and remuneration(${task.value.remuneration}) , continue?`, 'Undone Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'undonging task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgUndoneTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'undone task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("undone task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'undone task canceled', })
        })
    }

    const submitSuccessTask = async () => {
      if (task.value.employer != address.value && task.value.developer != address.value) {
        ElMessage({
          type: 'warning',
          message: 'You have no permission to success this task, connect keplr or switch to right employer or right developer',
        })
        return;
      }
      ElMessageBox.confirm(`if you are employer, it will send you remuneration(${task.value.remuneration}) to developer and send back deposit(${task.value.deposit}) for you; if you are a developer, it succeed only after your delivery exceeds the confirmation time of the employer! continue?`, 'Success Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'success task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgSuccessTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'success task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("cancel task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'cancel task canceled', })
        })
    }

    const submitFailTask = async () => {
      if (task.value.employer != address.value && task.value.developer != address.value) {
        ElMessage({
          type: 'warning',
          message: 'You have no permission to fail this task, connect keplr or switch to right employer or right developer',
        })
        return;
      }
      ElMessageBox.confirm(`if you are employer, it will send back you remuneration(${task.value.remuneration}) and deposit(${task.value.deposit}) and transfer developer's collateral(${task.value.collateral}) to you, it succeed only after developer delivery exceeds the deadline! if you are a developer, the system will send your collateral(${task.value.collateral}) to the employer! continue?`, 'Fail Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'fail task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgFailTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'fail task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("cancel task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'cancel task canceled', })
        })
    }

    const submitFeedbackTask = async () => {
      if (task.value.employer != address.value && task.value.developer != address.value) {
        ElMessage({
          type: 'warning',
          message: 'You have no permission to feedback this task, connect keplr or switch to right employer or right developer',
        })
        return;
      }

      ElMessageBox.prompt('you only have one chance to give feedback', 'Feedback Task', { confirmButtonText: 'OK', cancelButtonText: 'Cancel' })
        .then(async ({ value }) => {
          const loading = ElLoading.service({ lock: true, text: 'feedback task...' });
          try {
            const data = { creator: address.value, id: task.value.id, feedback: value }
            const reply = await $s.dispatch("sideline.sideline/sendMsgFeedbackTask", { value: data });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'feedback task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("cancel task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'feedback task canceled', })
        })
    }

    const submitStartJudgeTask = async () => {
      if (task.value.employer != address.value && task.value.developer != address.value) {
        ElMessage({
          type: 'warning',
          message: 'You have no permission to start judge this task, connect keplr or switch to right employer or right developer',
        })
        return;
      }
      ElMessageBox.confirm(`After start judge task, the validator will vote. the result for employer, win: it will send back you remuneration(${task.value.remuneration}) and deposit(${task.value.deposit}) and transfer developer's collateral(${task.value.collateral}) to you, lose: send you remuneration(${task.value.remuneration}) and deposit(${task.value.deposit}) to developer. the result for developer are the same! continue?`, 'Start Judge Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'start judge task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgStartJudgeTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'start judge task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("start judge task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'start judge task canceled', })
        })
    }

    const submitVoteTask = async () => {
      ElMessageBox.prompt('only validator can vote and only have one chance to vote, input "yes" support accuser, input "no" oppose accuser', 'Vote Task', { confirmButtonText: 'OK', cancelButtonText: 'Cancel' })
        .then(async ({ value }) => {
          let option = 0;
          if (value == "yes") {
            option = 0
          } else if (value == "no") {
            option = 1
          } else {
            ElMessage({ type: 'error', message: "Your input is incorrect" })
            return
          }

          const loading = ElLoading.service({ lock: true, text: 'vote task...' });
          try {
            const data = { creator: address.value, id: task.value.id, option }
            const reply = await $s.dispatch("sideline.sideline/sendMsgVoteTask", { value: data });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'vote task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("vote task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'vote task canceled', })
        })
    }
    const submitJudgeTask = async () => {
      ElMessageBox.confirm(`Excute judge results, continue?`, 'Judge Task', vnode)
        .then(async () => {
          const loading = ElLoading.service({ lock: true, text: 'judge task...' });
          try {
            const value = { creator: address.value, id: task.value.id }
            const reply = await $s.dispatch("sideline.sideline/sendMsgJudgeTask", { value });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({ type: 'success', message: 'judge task completed', })
              await getData()
            } else {
              ElMessage({ type: 'error', message: reply.rawLog, })
            }
          } catch (error) {
            console.log("judge task error", error)
            ElMessage({ type: 'error', message: error, })
          }
          loading.close()
        })
        .catch(() => {
          ElMessage({ type: 'info', message: 'judge task canceled', })
        })
    }

    return {
      blockHeight,
      msg,
      id,
      address,
      balance,
      validator,
      validatorBalance,
      sideline,
      sidelineBalance,
      developerBalance,
      employerBalance,
      preview,
      task,
      dialogVisible,
      rules,
      form,
      formRef,
      getData,
      initForm,
      closeDialog,
      submitCreateTask,
      formatTaskStatus,
      submitCancelTask,
      submitDoTask,
      submitSubmitTask,
      submitUndoneTask,
      submitSuccessTask,
      submitFailTask,
      submitFeedbackTask,
      submitStartJudgeTask,
      submitVoteTask,
      submitJudgeTask,
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
