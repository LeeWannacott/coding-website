
const audiobookReviews = [
  { id: 1, review: "A great book!", stars: 5 },
  { id: 2, review: "Average book", stars: 3 },
  { id: 3, review: "Worst book I have ever read", stars: 1 },
];
let sum = 0
for (let i = 0; i < audiobookReviews.length; i++) {
  sum += audiobookReviews[i].stars
}
console.log(sum);
