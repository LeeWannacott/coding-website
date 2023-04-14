const audiobookReviews = [
  { id: 1, review: "it was a good book", stars: 5 },
  { id: 2, review: "average book", stars: 3 },
  { id: 3, review: "worst book I have ever read", stars: 0 },
];
const reviews = [];
for (let i = 0; i < audiobookReviews.length; i++) {
  reviews.push(audiobookReviews[i].review);
}
console.log(reviews);
