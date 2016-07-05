package food_test

import (
	. "github.com/sleepypikachu/still-good-food/food"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Food", func() {
	var (
		testData      Recipe
		fetchedRecipe Recipe
	)

	BeforeEach(func() {
		testData = Recipe{
			Name:        "Gooseberry \u0026 custard pies",
			Ingredients: []string{"6 sheets filo pastry", "50g butter, melted", "2 tbsp golden caster sugar", "vanilla ice cream, to serve (optional)", "200g golden caster sugar", "1 vanilla pod, split and seeds scraped out", "500g gooseberries, topped and tailed", "300ml double cream", "300ml full-fat milk", "2 vanilla pods, split and seeds scraped out", "4 large egg yolks", "50g golden caster sugar", "20g plain flour", "20g cornflour"},
			Steps:       []string{"First, make the compote. Put the sugar and vanilla pod with seeds in a saucepan with 3 tbsp water and bring to the boil. Try to gently swirl rather than stir. Drop in the gooseberries and simmer for 20 mins or until half the liquid has bubbled off. Divide between four small pie dishes and put in the fridge to cool while you make the custard.", "Put the cream, milk and vanilla pods with seeds in a saucepan and bring just to the boil. Meanwhile, in a large bowl, whisk the egg yolks, sugar, flour and cornflour together. Pour the hot milk mixture over the egg mixture, whisking constantly. Rinse out the saucepan, then pour the custard mixture into the pan. Cook on a low heat, whisking, for 5-6 mins until thick. Pour through a sieve over a large jug and discard the vanilla pods, then pour the custard over the compote in the pie dishes. Leave to cool.", "Heat oven to 180C/160C fan/gas 4. To make the pie lids, lay one sheet of the filo on the work surface and brush with the melted butter, then put another sheet on top. Repeat with the rest of the sheets. On the final layer, scatter over the sugar, then cut the pastry into four rectangles. Lightly scrunch each rectangle and place loosely on top of each pie dish. Bake the pies in the oven for 20 mins or until the pastry is golden and caramelised. Remove from the oven and serve."},
			Yield:       "Serves 4",
			Difficulty:  "More effort",
			Preparation: "25 mins",
			Cook:        "50 mins",
			Nutrition:   NutritionInfo{Kcal: "1157", Fat: "61g", Saturates: "35g", Carbs: "", Sugars: "90g", Fibre: "6g", Protein: "14g", Salt: "1g"},
		}
		fetchedRecipe, _ = Scrape("http://www.bbcgoodfood.com/recipes/gooseberry-custard-pies")
	})

	Describe("Checking by field in Recipe", func() {
		Context("Obtain Name from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Name).Should(BeEquivalentTo(testData.Name))
			})
		})
		Context("Obtain Ingredients from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Ingredients).Should(BeEquivalentTo(testData.Ingredients))
			})
		})
		Context("Obtain Steps from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Steps).Should(BeEquivalentTo(testData.Steps))
			})
		})
		Context("Obtain Yield from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Yield).Should(BeEquivalentTo(testData.Yield))
			})
		})
		Context("Obtain Difficulty from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Difficulty).Should(BeEquivalentTo(testData.Difficulty))
			})
		})
		Context("Obtain Preparation from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Preparation).Should(BeEquivalentTo(testData.Preparation))
			})
		})
		Context("Obtain Cook from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Cook).Should(BeEquivalentTo(testData.Cook))
			})
		})
		Context("Obtain Yield from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Yield).Should(BeEquivalentTo(testData.Yield))
			})
		})
		Context("Obtain Nutrition from fetchedRecipe", func() {
			It("Should match with the test data", func() {
				Expect(fetchedRecipe.Nutrition).Should(BeEquivalentTo(testData.Nutrition))
			})
		})
	})
	Describe("Checking by field in Nutrition of a Recipe", func() {
		Context("Obtain Kcal from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Kcal).Should(BeEquivalentTo(testData.Nutrition.Kcal))
			})
		})
		Context("Obtain Fat from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Fat).Should(BeEquivalentTo(testData.Nutrition.Fat))
			})
		})
		Context("Obtain Saturates from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Saturates).Should(BeEquivalentTo(testData.Nutrition.Saturates))
			})
		})
		Context("Obtain Carbs from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Carbs).Should(BeEquivalentTo(testData.Nutrition.Carbs))
			})
		})
		Context("Obtain Sugars from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Sugars).Should(BeEquivalentTo(testData.Nutrition.Sugars))
			})
		})
		Context("Obtain Fibre from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Fibre).Should(BeEquivalentTo(testData.Nutrition.Fibre))
			})
		})
		Context("Obtain Protein from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Protein).Should(BeEquivalentTo(testData.Nutrition.Protein))
			})
		})
		Context("Obtain Salt from the Nutrition field of the fetchedRecipe", func() {
			It("Should match with test data", func() {
				Expect(fetchedRecipe.Nutrition.Salt).Should(BeEquivalentTo(testData.Nutrition.Salt))
			})
		})
	})
})
