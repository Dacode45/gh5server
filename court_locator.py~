import sys
import json

### COURT LOCATOR SCRIPT (court_locator.py)
###
### Author: Elliott Battle (09-13-2015)
### Event: GlobalHack V
### Team: Rush Hour 3
###
### Purpose: Process queries from a JSON database by taking geographic points (represented as a lattitude and a longitude)
###          and identifying which court in the desired county (if any) services that area.

# PATHING CONSTANTS (Hardcoded for now)
# A JSON database of geographic information for all of the municipalities in the county.
muni_filename = "stl_munis.json"
# A JSON database detailing location and contact information for all courts in the county.
court_filename = "stl_courts.json"
# A JSON database used to map together entries in the prior two databases based-on which courts
# service which areas.
service_map_filename = "stl_service_map.json"

def point_is_in_bounds(x, y, bounds):
    # POINT IS IN BOUNDS
    #
    # Checks if the point (x, y) is contained in the region bounded by the bounds variable, which is
    # an array of points defining the boundary of a polygon.
    #
    # Credit for Algorithm: http://www.ariel.com.au/a/python-point-int-poly.html
    n = len(bounds)
    inside = False
    p1x, p1y = bounds[0]
    for i in range(n+1):
        p2x,p2y = bounds[i % n]
        if y > min(p1y,p2y):
            if y <= max(p1y,p2y):
                if x <= max(p1x,p2x):
                    if p1y != p2y:
                        xinters = (y-p1y)*(p2x-p1x)/(p2y-p1y)+p1x
                    if p1x == p2x or x <= xinters:
                        inside = not inside
        p1x,p1y = p2x,p2y

    return inside

def point_is_in_poly(x, y, poly):
    # POINT IS IN POLYGON
    #
    # Checks if the point (x, y) is contained in a particular GEOJSON polygon.
    # Because this data is based-on the GEOJSON standard, the polygon object is a nested array of points.
    # Each inner array of points describes a closed curve. The first curve is the outter boundary of the polygon,
    # and all subsequent curves describe areas "cut-out" of the polygon, i.e. regions within the outermost curve
    # which are "holes" in the shape and therefore not considered part of the polygon.

    # Again, this comes first by GEOJSON convention. All other curves describe exclusion regions.
    outer_bound = poly[0]

    # We first check naively whether the point is even contained in the "whole" version of the polygon.
    if point_is_in_bounds(x, y, outer_bound):
        # Now we need to check if the point lies in a region excluded from the polygon.
        if len(poly) == 1:
            # The polygon has no holes and the point is inside its outer boundary. Polygon contains the point.
            return True
        else:
            # The polygon has some nonzero number of holes. Check to make sure the point isn't in any of them.
            for hole in poly[1:len(poly)]:
                if point_is_in_bounds(x, y, hole):
                    #If in a hole, it's not contained by the polygon. Return false immediately.
                    return False
            # The point is not excluded by any subsequent curves; it lies within the polygon.
            return True
    else:
        # Point is outside of the polygon's bounding region entirely. Return false.
        return False

def point_is_in_multipoly(x, y, multipoly):
    # POINT IS IN MULTIPOLYGON
    #
    # Checks if the point (x, y) is contained in a particular GEOJSON multipolygon.
    # A multipolygon is an array of regular GEOJSON polygons which may or may not be contigious when rendered in
    # a plane. A point is considered to be contained by a composite polygon if it is considered contained by any
    # one of the constituent polygons.

    # We simply iterate through the array, and return true if any are true, and false if all are false.
    for poly in multipoly:
        #print("Next poly has", len(poly), "points.")
        if point_is_in_poly(x, y, poly):
            return True
    return False

def muniForPoint(x, y):
    # MUNICIPALITY FOR POINT
    #
    # Determines which municipality (if any) within the county specified intersects lattitude x and longitude y.
    # While this method may return None, the driver and calling functions check for this sentinel condition.

    # Load the list of municipal boundaries as JSON objects (each being geometrically either a polygon or a multipolygon).
    # Save them into an array and iterate over it.
    file = open(muni_filename, 'r')
    muni_json_data = json.loads(file.read())
    file.close()
    munis = muni_json_data["munis"]

    # Iterate through those objects and look for one containing the point (x, y) within its bounds.
    for muni in munis:
        # Before we can iterate, we must know what type of polygon we're dealing with.
        muni_geometry = muni["geometry"]
        #print("Checking for match in:", muni["properties"]["name"])
        if muni_geometry["type"] == "Polygon":
            # Geometry is a polygon. Use basic containment algorithm.
            if point_is_in_poly(x, y, muni_geometry["coordinates"]):
                return muni
        elif muni_geometry["type"] == "MultiPolygon":
            # Geometry is a multipolygon.
            # Use the iterative containment algorithm over all member polygons.
            if point_is_in_multipoly(x, y, muni_geometry["coordinates"]):
                # Return if we have a match.
                return muni
        else:
            # Log that there was a geometry error, but don't stop iterating immediately;
            # the point could still be in a subsequent municipality and we don't want to miss it.
            # print("ERROR. INVALID GEOMETRY TYPE")

    # Return nothing if there's no match.
    return None

def courtForMuni(muni):
    # COURT FOR MUNICIPALITY
    #
    # Input: A municipality, as stored in a JSON object.
    # Output: Either the court presiding over that municipality or (if there is none or there's an unexpected error)
    #         the default option  county courts

    file = open(court_filename, 'r')
    court_json_data = json.loads(file.read())
    courts = court_json_data["courts"]
    file.close()

    if muni == None or muni["properties"]["muni_id"] == 0:
        # The if-statement above covers a few critical edge cases:
        # 1) The area is unincorporated and under county jurisdiction.
        # 2) The user inputs a point to which no municipality in the county corresponds
        #    (e.g. picks a point somewhere outside the county or even the state).
        # 3) The user picks a valid point within the county, but it is on the boundary between two municipalities.
        #    This is an edge case from the point containment algorithm which could have real-life consequences:
        #    tickets issued so close to a municipal boundary could be under ambiguous jurisdiction as-is.
        # 4) A point is in an unincorporated area without a name. This is caused by an oversight in the database.
        # 5) Some other general database error causes a valid municipality to not be found.

        # In each of these cases, we should return the county-level courthouse in this siutation (always at index 0) so
        # that the user can get some human assistance.
        return courts[0]
    else:
        # Otherwise, look-up the municipality in the service mapping table,
        # and return the court with the corresponding id.
        file = open(service_map_filename, 'r')
        service_map_json_data = json.loads(file.read())
        service_mappings = service_map_json_data["service_mappings"]
        file.close()

        # Iterate through the list of service mappings to find the
        # court_id servicing the municipality given.
        for s_map in service_mappings:
            if s_map["municipality"].lower() == muni["properties"]["name"].lower():
                court_id = s_map["servicing_court_id"]
                for court in courts:
                    if court_id == court["court_id"] or muni["properties"]["name"].lower() == court["city"].lower():
                        return court
                # A court serving a municipality with the same name was found in our mapping file,
                # but neither the court's municipality name nor id is in our list of courts.
                # This may mean  the JSON is incomplete or corrupt, so we have no choice but to return the default.
                return courts[0]
        # If we don't get a match anywhere, then that means we forgot to include
        # an incorporated municipality in our mapping JSON file.
        # Direct the user to the county-level offices for further assistance.
        return courts[0]

def main():
    # MAIN DRIVER ROUTINE:
    #
    # Takes the x and y coordinates the user wishes to query as arguments.
    # Send to the ostream via print() JSON objects for the municipality and
    # court relevant to the query.

##    FOR TESTING FUNCTIONS ONLY (DISABLE WHEN USING COMAND-LINE ARGS!!!)
##
##    print("Enter x, then y!")
##    x = eval(input())
##    y = eval(input())

    # Unpack the command-line arguments into lattitude x and longitude y
    x = sys.argv[0]
    y = sys.argv[1]

    # Find the municipality containing a point at lattitude x and longitude y
    muni = muniForPoint(x, y)

    # Find the court servicing that area. Default to the county court if one cannot be found.
    court = courtForMuni(muni)

    # Export the results of the query to the server-side software calling this script via a
    # print() statement
    print(json.dumps(court))

main()
