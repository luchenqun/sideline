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

export const taskColor = (status) => {
  if (status == "0") return "#337ecc";
  if (status == "1") return "#779649";
  if (status == "2") return "#6c8650";
  if (status == "3") return "#fff100";
  if (status == "4") return "#529b2e";
  if (status == "5") return "#e60012";
  if (status == "6") return "#8a1874";
  if (status == "7") return "#a29bfe";
  if (status == "8") return "#6c5ce7";
  if (status == "9") return "#1e2732";
  return "#000";
};

export const clone = (value) => JSON.parse(JSON.stringify(value));

export const sleep = (time) => {
  return new Promise(function (resolve) {
    setTimeout(resolve, time);
  });
};
