using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Migrations.Sqlite.Migrations
{
    /// <inheritdoc />
    public partial class Sqlite : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "FoodItems",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    Brand = table.Column<string>(type: "TEXT", nullable: false),
                    ProductName = table.Column<string>(type: "TEXT", nullable: false),
                    NutritionalContent_Protein = table.Column<double>(type: "REAL", nullable: false),
                    NutritionalContent_Carbohydrate = table.Column<double>(type: "REAL", nullable: false),
                    NutritionalContent_Fat = table.Column<double>(type: "REAL", nullable: false),
                    NutritionalContent_Kcal = table.Column<double>(type: "REAL", nullable: false),
                    OwnerId = table.Column<Guid>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_FoodItems", x => x.Id);
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "FoodItems");
        }
    }
}
