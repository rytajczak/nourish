import axios from "axios";

export default eventHandler(async () => {
  const headers = {
    "x-rapidapi-key": process.env.RAPIDAPI_KEY!,
    "x-rapidapi-host": process.env.RAPIDAPI_HOST!,
    "Content-Type": "application/json",
  };

  let res = await axios.get(
    `https://${process.env.RAPIDAPI_HOST}/recipes/mealplans/generate?timeFrame=week&targetCalories=2000`,
    { headers },
  );

  const daysOfWeek = [
    "monday",
    "tuesday",
    "wednesday",
    "thursday",
    "friday",
    "saturday",
    "sunday",
  ];
  const weekPlan = {};

  res.data.items.forEach((item) => {
    const dayName = daysOfWeek[item.day - 1];
    const meal = {
      id: JSON.parse(item.value).id,
      imageType: JSON.parse(item.value).imageType,
      title: JSON.parse(item.value).title,
      slot: item.slot,
      position: item.position,
      mealPlanId: item.mealPlanId,
      type: item.type,
    };

    if (!weekPlan[dayName]) {
      weekPlan[dayName] = [];
    }
    weekPlan[dayName].push(meal);
  });

  return weekPlan;
});
