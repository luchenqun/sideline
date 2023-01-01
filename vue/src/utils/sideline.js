export const taskStatusList = [
  {
    value: "0",
    label: "Created",
  },
  {
    value: "1",
    label: "Doing",
  },
  {
    value: "2",
    label: "Submited",
  },
  {
    value: "3",
    label: "Undone",
  },
  {
    value: "4",
    label: "Success",
  },
  {
    value: "5",
    label: "Fail",
  },
  {
    value: "6",
    label: "Judging",
  },
  {
    value: "7",
    label: "EmployerWin",
  },
  {
    value: "8",
    label: "DeveloperWin",
  },
  {
    value: "9",
    label: "Canced",
  },
];

export const formatTaskStatus = (value) => {
  for (const item of taskStatusList) {
    if (item.value == value) return item.label;
  }
  return `未知状态(${value})`;
};

export const clone = (value) => JSON.parse(JSON.stringify(value));

export const sleep = (time) => {
  return new Promise(function (resolve) {
    setTimeout(resolve, time);
  });
};
