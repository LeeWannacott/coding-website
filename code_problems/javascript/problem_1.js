const audiobookReviews = [
  { id: 1, review: "A great book!", stars: 5 },
  { id: 2, review: "Average book", stars: 3 },
  { id: 3, review: "Worst book I have ever read", stars: 1 },
];
const reviews = [];
for (let i = 0; i < audiobookReviews.length; i++) {
  reviews.push(audiobookReviews[i].review);
}
console.log(reviews);
