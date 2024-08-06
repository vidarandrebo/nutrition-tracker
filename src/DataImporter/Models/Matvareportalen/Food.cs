using Newtonsoft.Json;

namespace NutritionTracker.DataImporter.Models.Matvareportalen;

    public class Calories
    {
        [JsonProperty("sourceId")]
        public string SourceId { get; set; }

        [JsonProperty("quantity")]
        public double? Quantity { get; set; }

        [JsonProperty("unit")]
        public string Unit { get; set; }
    }

    public class Constituent
    {
        [JsonProperty("sourceId")]
        public string SourceId { get; set; }

        [JsonProperty("nutrientId")]
        public string NutrientId { get; set; }

        [JsonProperty("quantity")]
        public double? Quantity { get; set; }

        [JsonProperty("unit")]
        public string Unit { get; set; }
    }

    public class EdiblePart
    {
        [JsonProperty("percent")]
        public int Percent { get; set; }

        [JsonProperty("sourceId")]
        public string SourceId { get; set; }
    }

    public class Energy
    {
        [JsonProperty("sourceId")]
        public string SourceId { get; set; }

        [JsonProperty("quantity")]
        public double? Quantity { get; set; }

        [JsonProperty("unit")]
        public string Unit { get; set; }
    }

    public class Food
    {
        [JsonProperty("searchKeywords")]
        public List<string> SearchKeywords { get; set; }

        [JsonProperty("calories")]
        public Calories Calories { get; set; }

        [JsonProperty("portions")]
        public List<Portion> Portions { get; set; }

        [JsonProperty("ediblePart")]
        public EdiblePart EdiblePart { get; set; }

        [JsonProperty("langualCodes")]
        public List<string> LangualCodes { get; set; }

        [JsonProperty("energy")]
        public Energy Energy { get; set; }

        [JsonProperty("foodName")]
        public string FoodName { get; set; }

        [JsonProperty("latinName")]
        public string LatinName { get; set; }

        [JsonProperty("constituents")]
        public List<Constituent> Constituents { get; set; }

        [JsonProperty("uri")]
        public string Uri { get; set; }

        [JsonProperty("foodGroupId")]
        public string FoodGroupId { get; set; }

        [JsonProperty("foodId")]
        public string FoodId { get; set; }
    }

    public class Portion
    {
        [JsonProperty("portionName")]
        public string PortionName { get; set; }

        [JsonProperty("portionUnit")]
        public string PortionUnit { get; set; }

        [JsonProperty("quantity")]
        public double? Quantity { get; set; }

        [JsonProperty("unit")]
        public string Unit { get; set; }
    }

    public class Root
    {
        [JsonProperty("foods")]
        public List<Food> Foods { get; set; }

        [JsonProperty("locale")]
        public string Locale { get; set; }
    }

