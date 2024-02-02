using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Migrations.Postgres.Migrations
{
    /// <inheritdoc />
    public partial class Postgres : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "FoodItems",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uuid", nullable: false),
                    Brand = table.Column<string>(type: "text", nullable: false),
                    ProductName = table.Column<string>(type: "text", nullable: false),
                    NutritionalContent_Protein = table.Column<double>(type: "double precision", nullable: false),
                    NutritionalContent_Carbohydrate = table.Column<double>(type: "double precision", nullable: false),
                    NutritionalContent_Fat = table.Column<double>(type: "double precision", nullable: false),
                    NutritionalContent_Kcal = table.Column<double>(type: "double precision", nullable: false),
                    OwnerId = table.Column<Guid>(type: "uuid", nullable: false)
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
